// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"image/color"
	"image/draw"
	"math"
)

type pixel struct {
	r, g, b, a uint8
	average    uint8
	lightness  uint8
	luminosity uint8
}

// apply the given color mapping to the specified image buffers.
func apply(from, to *Rule, src, dst draw.Image) {
	var x, y int
	var clr color.Color

	rect := src.Bounds()

	for y = 0; y < rect.Dy(); y++ {
		for x = 0; x < rect.Dx(); x++ {
			clr = src.At(x, y)
			clr = transform(clr, from, to)

			dst.Set(x, y, clr)
		}
	}
}

// transform transforms the given color using the specified mapping.
// But only if it matches the filter rule.
func transform(clr color.Color, from, to *Rule) color.Color {
	var pix pixel

	r, g, b, a := clr.RGBA()
	pix.r, pix.g, pix.b, pix.a = uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8)

	if !match(&pix, from) {
		return clr
	}

	// Compute three different types of grayscale conversion.
	// These can be applied by named references.
	pix.average = uint8(((r + g + b) / 3) >> 8)
	pix.lightness = uint8(((min(min(r, g), b) + max(max(r, g), b)) / 2) >> 8)

	// For luminosity it is necessary to apply an inverse of the gamma
	// function for the color space before calculating the inner product.
	// Then you apply the gamma function to the reduced value. Failure to
	// incorporate the gamma function can result in errors of up to 20%.
	//
	// For typical computer stuff, the color space is sRGB. The right
	// numbers for sRGB are approx. 0.21, 0.72, 0.07. Gamma for sRGB
	// is a composite function that approximates exponentiation by 1/2.2
	//
	// This is a rather expensive operation, but gives a much more accurate
	// and satisfactory result than the average and lightness versions.
	pix.luminosity = gammaSRGB(0.212655*invGammaSRGB(pix.r) +
		0.715158*invGammaSRGB(pix.g) +
		0.072187*invGammaSRGB(pix.b))

	// Transform pixel.
	return color.RGBA{
		_transform(&pix, pix.r, to.R),
		_transform(&pix, pix.g, to.G),
		_transform(&pix, pix.b, to.B),
		_transform(&pix, pix.a, to.A),
	}
}

// _transform transforms a single channel using the specified mapping.
func _transform(pix *pixel, curr uint8, to Channel) uint8 {
	switch tt := to.(type) {
	case Number:
		switch tt.Operator {
		case "+":
			if tt.Percentage {
				perc := float64(curr) * 0.01
				v := int(curr) + int(perc*float64(tt.Value))
				if v > 0xff {
					v = 0xff
				}
				return uint8(v)
			}

			v := int(curr) + int(tt.Value)
			if v > 0xff {
				v = 0xff
			}
			return uint8(v)

		case "-":
			if tt.Percentage {
				perc := float64(curr) * 0.01
				v := int(curr) - int(perc*float64(tt.Value))
				if v < 0 {
					v = 0
				}
				return uint8(v)
			}

			v := int(curr) - int(tt.Value)
			if v < 0 {
				v = 0
			}
			return uint8(v)

		default:
			if tt.Percentage {
				perc := float64(curr) * 0.01
				v := int(perc * float64(tt.Value))
				if v > 0xff {
					v = 0xff
				}
				return uint8(v)
			}

			return tt.Value
		}

	case Name:
		switch tt {
		case NameR:
			return pix.r
		case NameG:
			return pix.g
		case NameB:
			return pix.b
		case NameA:
			return pix.a
		case NameLightness:
			return pix.lightness
		case NameLuminosity:
			return pix.luminosity
		case NameAverage:
			return pix.average
		}
	}

	// wildcard -- return original value.
	return curr
}

// match checks of the given color matches the given filter.
func match(pix *pixel, filter *Rule) bool {
	return matchChannel(pix.r, filter.R) &&
		matchChannel(pix.g, filter.G) &&
		matchChannel(pix.b, filter.B) &&
		matchChannel(pix.a, filter.A)
}

func matchChannel(v uint8, c Channel) bool {
	num, ok := c.(Number)
	if !ok {
		return true // wildcard
	}

	switch num.Operator {
	case "<":
		return v < num.Value
	case "<=":
		return v <= num.Value
	case ">":
		return v > num.Value
	case ">=":
		return v >= num.Value
	}

	return true
}

// sRGB "gamma" function (approx 2.2)
func gammaSRGB(v float64) uint8 {
	if v <= 0.0031308 {
		v *= 12.92
	} else {
		v = 1.055*math.Pow(v, 1.0/2.4) - 0.055
	}

	return uint8(v*255 + .5)
}

// Inverse of gamma_sRGB "gamma" function. (approx 2.2)
func invGammaSRGB(ic uint8) float64 {
	c := float64(ic) / 255
	if c <= 0.04045 {
		return c / 12.92
	}

	return math.Pow(((c + 0.055) / (1.055)), 2.4)
}

func min(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}

func max(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

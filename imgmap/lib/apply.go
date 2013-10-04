// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

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

// Apply the given color mapping to the specified image buffers.
func Apply(from, to *Rule, src, dst draw.Image) {
	var x, y int
	var r, g, b, a uint32
	var sc color.Color
	var dc color.RGBA
	var pix pixel

	rect := src.Bounds()

	for y = 0; y < rect.Dy(); y++ {
		for x = 0; x < rect.Dx(); x++ {
			sc = src.At(x, y)

			r, g, b, a = sc.RGBA()
			pix.r = uint8(r >> 8)
			pix.g = uint8(g >> 8)
			pix.b = uint8(b >> 8)
			pix.a = uint8(a >> 8)

			// Check if the pixel matches the filter rule.
			if !(match(pix.r, from.R) && match(pix.g, from.G) && match(pix.b, from.B) && match(pix.a, from.A)) {
				dst.Set(x, y, sc)
				continue
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
			pix.luminosity = gammaSRGB(
				0.212655*invGammaSRGB(pix.r) +
					0.715158*invGammaSRGB(pix.g) +
					0.072187*invGammaSRGB(pix.b))

			// Transform color.
			dc.R = transform(&pix, pix.r, to.R)
			dc.G = transform(&pix, pix.g, to.G)
			dc.B = transform(&pix, pix.b, to.B)
			dc.A = transform(&pix, pix.a, to.A)

			// Set new pixel.
			dst.Set(x, y, dc)
		}
	}
}

// transform transforms a single channel using the specified mapping.
func transform(pix *pixel, curr uint8, to Channel) uint8 {
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

// match checks if the given channel value mathes the given channel rule.
func match(v uint8, c Channel) bool {
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

	return v == num.Value
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

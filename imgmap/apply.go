// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"image/color"
	"image/draw"
	"math"
)

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
	r, g, b, a := clr.RGBA()
	br, bg, bb, ba := uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8)

	if !match(br, bg, bb, ba, from) {
		return clr
	}

	// Compute three different types of greyscale conversion.
	// These can be applied by named references.
	average := uint8(((r + g + b) / 3) >> 8)
	lightness := uint8(((min(min(r, g), b) + max(max(r, g), b)) / 2) >> 8)
	luminosity := gammaSRGB(0.212655*invGammaSRGB(br) + 0.715158*invGammaSRGB(bg) + 0.072187*invGammaSRGB(bb))

	// Transform pixel.
	return color.RGBA{
		_transform(br, bg, bb, ba, br, average, lightness, luminosity, to.R),
		_transform(br, bg, bb, ba, bg, average, lightness, luminosity, to.G),
		_transform(br, bg, bb, ba, bb, average, lightness, luminosity, to.B),
		_transform(br, bg, bb, ba, ba, average, lightness, luminosity, to.A),
	}
}

// _transform transforms a single channel using the specified mapping.
func _transform(sr, sg, sb, sa, curr, average, lightness, luminosity uint8, to Channel) uint8 {
	switch tt := to.(type) {
	case Number:
		switch tt.Operator {
		case "+":
			if tt.Percentage {
				perc := float64(curr>>8) * 0.01
				v := int(curr) + int(perc*float64(tt.Value))
				return uint8(v % 0xff)
			}

			v := int(curr) + int(tt.Value)
			return uint8(v % 0xff)

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
				return uint8(v % 0xff)
			}

			return tt.Value
		}

	case Name:
		switch tt {
		case NameR:
			return sr
		case NameG:
			return sg
		case NameB:
			return sb
		case NameA:
			return sa
		case NameLightness:
			return lightness
		case NameLuminosity:
			return luminosity
		case NameAverage:
			return average
		}
	}

	// wildcard -- return original value.
	return curr
}

// match checks of the given color matches the given filter.
func match(r, g, b, a uint8, filter *Rule) bool {
	return matchChannel(r, filter.R) && matchChannel(g, filter.G) &&
		matchChannel(b, filter.B) && matchChannel(a, filter.A)
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

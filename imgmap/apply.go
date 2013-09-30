// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"image/color"
	"image/draw"
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

	if !match(uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8), from) {
		return clr
	}

	mean := (r + g + b) / 3

	return color.RGBA{
		_transform(r, g, b, a, r, mean, to.R),
		_transform(r, g, b, a, g, mean, to.G),
		_transform(r, g, b, a, b, mean, to.B),
		_transform(r, g, b, a, a, mean, to.A),
	}
}

// _transform transforms a single channel using the specified mapping.
func _transform(sr, sg, sb, sa, curr, mean uint32, to Channel) uint8 {
	switch tt := to.(type) {
	case Number:
		switch tt.Operator {
		case "+":
			if tt.Percentage {
				perc := float64(curr>>8) * 0.01
				v := int(curr>>8) + int(perc*float64(tt.Value))
				return uint8(v % 0xff)
			}

			v := int(curr>>8) + int(tt.Value)
			return uint8(v % 0xff)

		case "-":
			if tt.Percentage {
				perc := float64(curr>>8) * 0.01
				v := int(curr>>8) - int(perc*float64(tt.Value))
				if v < 0 {
					v = 0
				}
				return uint8(v)
			}

			v := int(curr>>8) - int(tt.Value)
			if v < 0 {
				v = 0
			}
			return uint8(v)

		default:
			if tt.Percentage {
				perc := float64(curr>>8) * 0.01
				v := int(perc * float64(tt.Value))
				return uint8(v % 0xff)
			}

			return tt.Value
		}

	case Name:
		switch tt {
		case NameR:
			return uint8(sr >> 8)
		case NameG:
			return uint8(sg >> 8)
		case NameB:
			return uint8(sb >> 8)
		case NameA:
			return uint8(sa >> 8)
		case NameMean:
			return uint8(mean)
		}
	}

	// wildcard -- return original value.
	return uint8(curr >> 8)
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

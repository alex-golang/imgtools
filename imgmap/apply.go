// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"image/color"
	"image/draw"
)

// apply the given color mapping to the specified image buffers.
func apply(from, to *Color, src, dst draw.Image) {
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
func transform(clr color.Color, from, to *Color) color.Color {
	if !match(clr, from) {
		return clr
	}

	r, g, b, a := clr.RGBA()
	mrgba := (r + g + b + a) / 4
	mrgb := (r + g + b) / 3

	return color.RGBA{
		_transform(r, g, b, a, r, mrgb, mrgba, to.R),
		_transform(r, g, b, a, g, mrgb, mrgba, to.G),
		_transform(r, g, b, a, b, mrgb, mrgba, to.B),
		_transform(r, g, b, a, a, mrgb, mrgba, to.A),
	}
}

// _transform transforms a single channel using the specified mapping.
func _transform(sr, sg, sb, sa, curr, mrgb, mrgba uint32, to Channel) uint8 {
	switch tt := to.(type) {
	case Number:

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
		case NameMeanRGBA:
			return uint8(mrgba)
		case NameMeanRGB:
			return uint8(mrgb)
		}
	}

	return uint8(curr >> 8)
}

// match checks of the given color matches the given filter.
func match(clr color.Color, filter *Color) bool {
	return true
}

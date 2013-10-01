// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"image"
)

// grayscale turns the image into a grayscale image.
func grayscale(img image.Image) image.Image {
	rect := img.Bounds()
	gray := image.NewGray(rect)

	var x, y int
	for y = rect.Min.Y; y < rect.Max.Y; y++ {
		for x = rect.Min.X; x < rect.Max.X; x++ {
			gray.Set(x, y, img.At(x, y))
		}
	}

	return gray
}

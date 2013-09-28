// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import (
	"image"
	"io"
)

// Decode decodes an image from the given stream.
// It returns the image and the name of the image's format.
func Decode(r io.Reader) (image.Image, string, error) {
	return image.Decode(r)
}

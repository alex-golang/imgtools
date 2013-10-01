// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import (
	"image"
	"image/jpeg"
	"io"
)

func init() {
	RegisterExtensions(".jpg", ".jpeg")
	RegisterEncoder("jpeg", encodeJpeg, "quality")
}

func encodeJpeg(w io.Writer, m image.Image, options OptionSet) error {
	return jpeg.Encode(w, m, &jpeg.Options{
		Quality: options.Int("quality", jpeg.DefaultQuality),
	})
}

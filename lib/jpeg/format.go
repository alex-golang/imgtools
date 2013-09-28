// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package jpeg

import (
	"github.com/jteeuwen/imgtools/lib"
	"image"
	"image/jpeg"
	"io"
)

func init() {
	lib.RegisterFormat("jpeg", encode, "quality")
}

func encode(w io.Writer, m image.Image, options lib.OptionSet) error {
	return jpeg.Encode(w, m, &jpeg.Options{
		Quality: options.Int("quality", jpeg.DefaultQuality),
	})
}

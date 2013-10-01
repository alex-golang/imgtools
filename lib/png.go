// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import (
	"image"
	"image/png"
	"io"
)

func init() {
	RegisterExtensions(".png")
	RegisterEncoder("png", func(w io.Writer, m image.Image, options OptionSet) error {
		return png.Encode(w, m)
	})
}

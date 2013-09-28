// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package png

import (
	"github.com/jteeuwen/imgtools/lib"
	"image"
	"image/png"
	"io"
)

func init() {
	lib.RegisterEncoder("png", encode)
}

func encode(w io.Writer, m image.Image, options lib.OptionSet) error {
	return png.Encode(w, m)
}

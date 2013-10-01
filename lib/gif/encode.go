// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package gif

import (
	"github.com/jteeuwen/imgtools/lib"
	"image"
	"io"
)

func init() {
	lib.RegisterExtensions(".gif")
	lib.RegisterEncoder("gif", encode, "quantizer")
}

func encode(w io.Writer, m image.Image, options lib.OptionSet) error {
	return nil
}

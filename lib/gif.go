// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import (
	"image"
	"io"
)

func init() {
	RegisterExtensions(".gif")
	RegisterEncoder("gif", func(w io.Writer, m image.Image, options OptionSet) error {
		return nil
	}, "quantizer")
}

// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import (
	"fmt"
	"image"
	"io"
	"strings"
)

// Encode encodes the given image and writes it to the specified stream.
// The format specifies which output format is desired. This must be
// one of the currently registered formats. The options string holds
// a semi-colon-separated list of key/value pairs with encoder options.
func Encode(w io.Writer, format string, m image.Image, options string) error {
	for _, enc := range formats {
		if !strings.EqualFold(format, enc.Name) {
			continue
		}

		enc.Options.Parse(options)
		return enc.Encode(w, m, enc.Options)
	}

	return fmt.Errorf("unsupported image format: %s", format)
}

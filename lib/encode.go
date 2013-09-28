// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import (
	"fmt"
	"image"
	"io"
	"strings"
)

type EncodeFunc func(io.Writer, image.Image, OptionSet) error

// List of registered encoders.
var Encoders []*Encoder

// RegisterEncoder registers the encoder for a given image format.
//
// The option list holds supported encoder option keys.
func RegisterEncoder(format string, ef EncodeFunc, options ...string) {
	Encoders = append(Encoders, &Encoder{
		Name:    format,
		Encode:  ef,
		Options: NewOptionSet(options...),
	})
}

// Encoder describes the encoder of a supported image type.
type Encoder struct {
	Name    string     // Name of the format: png, gif, pnm, etc
	Encode  EncodeFunc // Encode handler
	Options OptionSet  // Encoder options
}

// Encode encodes the given image and writes it to the specified stream.
// The format specifies which output format is desired. This must be
// one of the currently registered formats. The options string holds
// a semi-colon-separated list of key/value pairs with encoder options.
func Encode(w io.Writer, format string, m image.Image, options string) error {
	for _, enc := range Encoders {
		if !strings.EqualFold(format, enc.Name) {
			continue
		}

		enc.Options.Parse(options)
		return enc.Encode(w, m, enc.Options)
	}

	return fmt.Errorf("unsupported image format: %s", format)
}

// Keys returns the list of encoder option keys.
func (e *Encoder) Keys() []string {
	list := make([]string, 0, len(e.Options))

	for key := range e.Options {
		list = append(list, key)
	}

	return list
}

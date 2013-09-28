// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import (
	"image"
	"io"
	"strings"
)

type EncodeFunc func(io.Writer, image.Image, OptionSet) error

var formats []*Format

// RegisterFormat registers the encoder and decoder for
// a given image format.
//
// The option list holds supported encoder option keys.
func RegisterFormat(format string, ef EncodeFunc, options ...string) {
	formats = append(formats, &Format{
		Name:    format,
		Encode:  ef,
		Options: NewOptionSet(options...),
	})
}

// Format describes the decoder and encoder of a supported
// image type.
type Format struct {
	Name    string     // Name of the format: png, gif, pnm, etc
	Encode  EncodeFunc // Encode handler
	Options OptionSet  // Encoder options
}

// Formats returns a list of registered image format names.
func Formats() []string {
	list := make([]string, 0, len(formats))

	for _, f := range formats {
		list = append(list, f.Name)
	}

	return list
}

// Supported returns true if the given image format name
// is supported by this library.
func Supported(format string) bool {
	for _, f := range formats {
		if strings.EqualFold(format, f.Name) {
			return true
		}
	}
	return false
}

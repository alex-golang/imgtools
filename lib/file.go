// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import (
	"path"
	"strings"
)

// List of registered file extensions.
var extensions []string

// RegisterExtensions registers the given file extensions.
// The extensions are expected to be in the format: ".ext"
func RegisterExtensions(ext ...string) {
	extensions = append(extensions, ext...)
}

// ValidFile returns true if the given file path has a known
// file extension.
func ValidFile(file string) bool {
	ext := path.Ext(file)
	for _, v := range extensions {
		if strings.EqualFold(v, ext) {
			return true
		}
	}
	return false
}

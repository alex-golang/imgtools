// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import "strings"

// Encoders returns a list of registered image format names.
func Formats() []string {
	list := make([]string, 0, len(Encoders))

	for _, f := range Encoders {
		list = append(list, f.Name)
	}

	return list
}

// Supported returns true if the given image format name
// is supported by this library.
func Supported(format string) bool {
	for _, f := range Encoders {
		if strings.EqualFold(format, f.Name) {
			return true
		}
	}
	return false
}

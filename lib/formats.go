// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import "strings"

// Formats returns a list of supported image formats.
func Formats() []string {
	return []string {"png", "jpeg", "gif", "pnm"}
}

// Supported returns true if the given image format name
// is supported by this library.
func Supported(format string) bool {
	for _, v := range Formats() {
		if strings.EqualFold(format, v) {
			return true
		}
	}
	return false
}

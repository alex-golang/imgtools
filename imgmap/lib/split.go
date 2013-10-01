// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import (
	"bytes"
	"strings"
)

// splitString splits the input data and returns a list with empty elements removed.
func splitString(data, token string) []string {
	a := strings.Split(data, token)
	b := make([]string, 0, len(a))

	for _, v := range a {
		v = strings.TrimSpace(v)
		if len(v) > 0 {
			b = append(b, v)
		}
	}

	return b
}

// splitBytes splits the input data and returns a list with empty elements removed.
func splitBytes(data, token []byte) [][]byte {
	a := bytes.Split(data, token)
	b := make([][]byte, 0, len(a))

	for _, v := range a {
		v = bytes.TrimSpace(v)
		if len(v) > 0 {
			b = append(b, v)
		}
	}

	return b
}

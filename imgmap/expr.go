// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"image/draw"
	"os"
)

// parseExpression parses a single color map expression and
// applies it to the source and destination images.
func parseExpression(linecount int, line []byte, src, dst draw.Image) bool {
	// Strip whitespace and code comments.
	if idx := bytes.Index(line, semicolon); idx > -1 {
		line = line[:idx]
	}

	line = bytes.Replace(line, tab, space, -1)

	// Split into source and destination color mappings.
	list := splitBytes(line, space)
	if len(list) != 8 {
		return false
	}

	left := bytes.Join(list[:4], space)
	right := bytes.Join(list[4:], space)

	from, err := ParseColor(string(left))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Source color at line %d: %v", linecount, err)
		return false
	}

	to, err := ParseColor(string(right))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Destination color at line %d: %v", linecount, err)
		return false
	}

	if !validColor(from) {
		fmt.Fprintf(os.Stderr, "Invalid source color at line %d", linecount)
		return false
	}

	apply(from, to, src, dst)
	return true
}

// validColor returns true if the given color contains no
// invalid filter data for a source color.
func validColor(c *Color) bool {
	return validChannel(c.R) && validChannel(c.G) &&
		validChannel(c.B) && validChannel(c.A)
}

// validFrom returns true if the given color contains no
// invalid filter data for a source color.
func validChannel(c Channel) bool {

	return true
}

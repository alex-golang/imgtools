// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"fmt"
)

// Channel represents a color channel value.
// This can be a number, wildcard or named reference.
type Channel string

// Color represents a single rgba color.
type Color struct {
	R Channel
	G Channel
	B Channel
	A Channel
}

func (c *Color) String() string {
	return fmt.Sprintf("%s %s %s %s", c.R, c.G, c.B, c.A)
}

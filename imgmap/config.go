// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

// Config defines processed command line flags.
type Config struct {
	Input      string
	MapFile    string
	Monochrome *Color
	Saturate   Channel
	Brightness Channel
	Grayscale  bool
	BlackWhite bool
}

// Pipe returns true if the input filename is not defined.
// This means we should expect input to come through stdin.
func (c *Config) Pipe() bool {
	return len(c.Input) == 0
}

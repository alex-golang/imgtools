// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"fmt"
	"strings"
)

// Color represents a single rgba color.
type Color struct {
	R Channel
	G Channel
	B Channel
	A Channel
}

// ParseColor parses a color from the given input string.
func ParseColor(data string) (*Color, error) {
	data = strings.TrimSpace(data)
	if len(data) == 0 {
		return nil, fmt.Errorf("Invalid color value")
	}

	list := split(strings.Replace(data, " ", ",", -1), ",")

	if len(list) != 4 {
		return nil, fmt.Errorf("Invalid color value %q", data)
	}

	var err error
	c := new(Color)

	c.R, err = parseChannel(list[0])
	if err != nil {
		return nil, err
	}

	c.G, err = parseChannel(list[1])
	if err != nil {
		return nil, err
	}

	c.B, err = parseChannel(list[2])
	if err != nil {
		return nil, err
	}

	c.A, err = parseChannel(list[3])
	if err != nil {
		return nil, err
	}

	return c, nil
}

// split splits the input data and returns a list with empty elements removed.
func split(data, token string) []string {
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

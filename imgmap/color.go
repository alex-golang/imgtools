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

	switch len(list) {
	case 1:
		return parseHexColor(list[0])
	case 4:
		return parseRGBAColor(list[0], list[1], list[2], list[3])
	default:
		return nil, fmt.Errorf("Invalid color value %q", data)
	}
}

// parseHexColor parses a single hex value into a color.
// This should look like: 16#rrggbbaa
func parseHexColor(clr string) (*Color, error) {
	var err error

	if !strings.HasPrefix(clr, "0x") || len(clr) != 10 {
		return nil, fmt.Errorf("Invalid hexadecimal color value: %s", clr)
	}

	clr = clr[2:]

	c := new(Color)
	c.R, err = parseChannel("0x" + clr[:2])
	if err != nil {
		return nil, err
	}

	c.G, err = parseChannel("0x" + clr[2:4])
	if err != nil {
		return nil, err
	}

	c.B, err = parseChannel("0x" + clr[4:6])
	if err != nil {
		return nil, err
	}

	c.A, err = parseChannel("0x" + clr[6:])
	if err != nil {
		return nil, err
	}

	return c, nil
}

// parseRGBAColor parses 4 RGBA components into a color.
func parseRGBAColor(r, g, b, a string) (*Color, error) {
	var err error
	c := new(Color)

	c.R, err = parseChannel(r)
	if err != nil {
		return nil, err
	}

	c.G, err = parseChannel(g)
	if err != nil {
		return nil, err
	}

	c.B, err = parseChannel(b)
	if err != nil {
		return nil, err
	}

	c.A, err = parseChannel(a)
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

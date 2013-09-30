// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"fmt"
	"strings"
)

// Rule represents a single rgba Rule.
type Rule struct {
	R Channel
	G Channel
	B Channel
	A Channel
}

// ParseRule parses a Rule from the given input string.
func ParseRule(data string) (*Rule, error) {
	data = strings.TrimSpace(data)
	if len(data) == 0 {
		return nil, fmt.Errorf("Invalid Rule value")
	}

	list := splitString(strings.Replace(data, " ", ",", -1), ",")

	if len(list) != 4 {
		return nil, fmt.Errorf("Invalid Rule value %q", data)
	}

	var err error
	c := new(Rule)

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

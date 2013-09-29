// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"fmt"
	"strconv"
)

// Channel represents a color channel value.
// This can be a number, wildcard or named reference.
type Channel interface{}

// parseChannel parses a channel value from the given input string.
func parseChannel(data string) (Channel, error) {
	var sign uint8
	var percentage bool

	base := 10

	if len(data) > 2 && data[0] == '0' && data[1] == 'x' {
		base = 16
		data = data[2:]
	}

	switch {
	case data[0] == '?':
		return Wildcard{}, nil

	case len(data) == 2 && data[0] == '#':
		switch data[1] {
		case 'r', 'R':
			return NameR, nil
		case 'g', 'G':
			return NameG, nil
		case 'b', 'B':
			return NameB, nil
		case 'a', 'A':
			return NameA, nil
		default:
			return nil, fmt.Errorf("Invalid named reference: %s", data)
		}

	case data[0] == '-' || data[0] == '+':
		sign = data[0]
		data = data[1:]
	}

	if data[len(data)-1] == '%' {
		percentage = true
		data = data[:len(data)-1]
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("Invalid channel value: %s", data)
	}

	n, err := strconv.ParseUint(data, base, 8)
	if err != nil {
		return nil, fmt.Errorf("ParseChannel: %v", err)
	}

	return Number{
		Value:      uint8(n),
		Sign:       sign,
		Percentage: percentage,
	}, nil
}

// Wildcard represents a channel as a wildcard value.
type Wildcard struct{}

// Name represents a named channel reference.
type Name uint8

// Known names.
const (
	NameR Name = iota
	NameG
	NameB
	NameA
)

// Number represents a numeric value.
// It can cary a sign and/or percentile token.
type Number struct {
	Value      uint8
	Sign       uint8 // '+', '-' or 0
	Percentage bool
}

// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Channel represents a color channel value.
// This can be a number, wildcard or named reference.
type Channel interface{}

// parseChannel parses a channel value from the given input string.
func parseChannel(data string) (Channel, error) {
	var operator string
	var percentage bool

	base := 10

	switch {
	case strings.HasPrefix(data, "-"):
		operator = "-"
		data = data[1:]
	case strings.HasPrefix(data, "+"):
		operator = "+"
		data = data[1:]
	case strings.HasPrefix(data, ">="):
		operator = ">="
		data = data[2:]
	case strings.HasPrefix(data, ">"):
		operator = ">"
		data = data[1:]
	case strings.HasPrefix(data, "<="):
		operator = "<="
		data = data[2:]
	case strings.HasPrefix(data, "<"):
		operator = "<"
		data = data[1:]
	}

	if len(data) > 2 && strings.HasPrefix(data, "0x") {
		base = 16
		data = data[2:]
	}

	switch {
	case data[0] == '?':
		return Wildcard{}, nil

	case len(data) == 2 && data[0] == '#':
		switch data[1] {
		case 'r':
			return NameR, nil
		case 'g':
			return NameG, nil
		case 'b':
			return NameB, nil
		case 'a':
			return NameA, nil
		case 'm':
			return NameMean, nil
		default:
			return nil, fmt.Errorf("Invalid named reference: %s", data)
		}
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
		Operator:   operator,
		Value:      uint8(n),
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
	NameMean
)

// Number represents a numeric value.
// It can carry an operator and/or percentile token.
type Number struct {
	Operator   string // 0, +, -, <, <=, >, >=
	Value      uint8
	Percentage bool
}

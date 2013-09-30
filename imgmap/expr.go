// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"image/draw"
	"os"
)

// parseExpression parses a single Rule map expression and
// applies it to the source and destination images.
func parseExpression(line int, expr []byte, src, dst draw.Image) bool {
	// Strip whitespace and code comments.
	if idx := bytes.Index(expr, semicolon); idx > -1 {
		expr = expr[:idx]
	}

	expr = bytes.Replace(expr, tab, space, -1)

	// Split into source and destination Rule mappings.
	list := splitBytes(expr, space)
	if len(list) != 8 {
		return false
	}

	left := bytes.Join(list[:4], space)
	right := bytes.Join(list[4:], space)

	from, err := ParseRule(string(left))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Source Rule at line %d: %v\n", line, err)
		return false
	}

	to, err := ParseRule(string(right))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Destination Rule at line %d: %v\n", line, err)
		return false
	}

	if !validFrom(line, from) || !validTo(line, to) {
		return false
	}

	apply(from, to, src, dst)
	return true
}

// validFrom returns false if the given Rule does not contain values
// which make no sense in this context.
func validFrom(line int, r *Rule) bool {
	return validFromChannel(line, r.R) && validFromChannel(line, r.G) &&
		validFromChannel(line, r.B) && validFromChannel(line, r.A)
}

func validFromChannel(line int, c Channel) bool {
	switch tt := c.(type) {
	case Number:
		switch tt.Operator {
		case "", "<", "<=", ">", ">=":
		default:
			fmt.Fprintf(os.Stderr, "Line %d: Operator %q is not valid in a source Rule context.\n", line, tt.Operator)
			return false
		}

		if tt.Percentage {
			fmt.Fprintf(os.Stderr, "Line %d: Percentages are not valid in a source Rule context.\n", line)
			return false
		}

		return true

	case Name:
		fmt.Fprintf(os.Stderr, "Line %d: Named reference is not valid in a source Rule context.\n", line)
		return false
	}

	// wildcard
	return true
}

// validTo returns false if the given Rule does not contain values
// which make no sense in this context.
func validTo(line int, r *Rule) bool {
	return validToChannel(line, r.R) && validToChannel(line, r.G) &&
		validToChannel(line, r.B) && validToChannel(line, r.A)
}

func validToChannel(line int, c Channel) bool {
	num, ok := c.(Number)
	if ok {
		switch num.Operator {
		case "", "+", "-":
		default:
			fmt.Fprintf(os.Stderr, "Line %d: Operator %q is not valid in a destination Rule context.\n", line, num.Operator)
			return false
		}
	}

	// wildcard & named reference
	return true
}

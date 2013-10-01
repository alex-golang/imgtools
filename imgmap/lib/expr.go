// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import (
	"bytes"
	"fmt"
	"image/draw"
)

var (
	semicolon = []byte{';'}
	tab       = []byte{'\t'}
	space     = []byte{' '}
)

// Parse parses a single color map expression and
// applies it to the source and destination images.
func Parse(expr []byte, src, dst draw.Image) (bool, error) {
	// Strip whitespace and code comments.
	if idx := bytes.Index(expr, semicolon); idx > -1 {
		expr = expr[:idx]
	}

	expr = bytes.Replace(expr, tab, space, -1)

	// Split into source and destination Rule mappings.
	list := splitBytes(expr, space)
	if len(list) != 8 {
		return false, nil
	}

	left := bytes.Join(list[:4], space)
	right := bytes.Join(list[4:], space)

	from, err := ParseRule(string(left))
	if err != nil {
		return false, err
	}

	to, err := ParseRule(string(right))
	if err != nil {
		return false, err
	}

	err = validFrom(from)
	if err != nil {
		return false, err
	}

	err = validTo(to)
	if err != nil {
		return false, err
	}

	Apply(from, to, src, dst)
	return true, nil
}

// validFrom returns false if the given Rule does not contain values
// which make no sense in this context.
func validFrom(r *Rule) error {
	if err := validFromChannel(r.R); err != nil {
		return err
	}

	if err := validFromChannel(r.G); err != nil {
		return err
	}

	if err := validFromChannel(r.B); err != nil {
		return err
	}

	return validFromChannel(r.A)
}

func validFromChannel(c Channel) error {
	switch tt := c.(type) {
	case Number:
		switch tt.Operator {
		case "", "<", "<=", ">", ">=":
		default:
			return fmt.Errorf("Operator %q is not valid in a source Rule context.", tt.Operator)
		}

		if tt.Percentage {
			return fmt.Errorf("Percentages are not valid in a source Rule context.")
		}

	case Name:
		return fmt.Errorf("Named reference is not valid in a source Rule context.")
	}

	return nil
}

// validTo returns false if the given Rule does not contain values
// which make no sense in this context.
func validTo(r *Rule) error {
	if err := validToChannel(r.R); err != nil {
		return err
	}

	if err := validToChannel(r.G); err != nil {
		return err
	}

	if err := validToChannel(r.B); err != nil {
		return err
	}

	return validToChannel(r.A)
}

func validToChannel(c Channel) error {
	num, ok := c.(Number)
	if ok {
		switch num.Operator {
		case "", "+", "-":
		default:
			return fmt.Errorf("Operator %q is not valid in a destination Rule context.", num.Operator)
		}
	}

	// wildcard & named reference
	return nil
}

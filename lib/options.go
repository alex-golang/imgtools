// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import (
	"strconv"
	"strings"
)

// OptionSet is a set of key/value pairs,
// intended to supply encoding OptionSet to the various
// image format encoders.
//
// While keys and values are provided as strings,
// this type provides convenience methods to convert
// the value to various basic data types.
type OptionSet map[string]string

// NewOptionSet creates a new, empty encoder Option set.
// The specified keys list defines supported key names.
func NewOptionSet(keys ...string) OptionSet {
	o := make(OptionSet)

	for _, key := range keys {
		o[key] = ""
	}

	return o
}

// Parse parses the input string as a set of
// key/value pairs. For example:
//
//    quality:100;width:640;height:480
//
func (o OptionSet) Parse(data string) {
	data = strings.TrimSpace(data)
	if len(data) == 0 {
		return
	}

	pairs := strings.Split(data, ";")

	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if len(pair) == 0 {
			continue
		}

		set := strings.Split(pair, ":")
		if len(set) != 2 {
			continue
		}

		key := strings.TrimSpace(set[0])
		value := strings.TrimSpace(set[1])

		if len(key) == 0 || len(value) == 0 {
			continue
		}

		o[key] = value
	}
}

// Int returns the value for the given key as int.
func (o OptionSet) Int(key string, defaultval int) int {
	return int(o.Int64(key, int64(defaultval)))
}

// Int8 returns the value for the given key as int8.
func (o OptionSet) Int8(key string, defaultval int8) int8 {
	return int8(o.Int64(key, int64(defaultval)))
}

// Int16 returns the value for the given key as int16.
func (o OptionSet) Int16(key string, defaultval int16) int16 {
	return int16(o.Int64(key, int64(defaultval)))
}

// Int32 returns the value for the given key as int32.
func (o OptionSet) Int32(key string, defaultval int32) int32 {
	return int32(o.Int64(key, int64(defaultval)))
}

// Int64 returns the value for the given key as int64.
func (o OptionSet) Int64(key string, defaultval int64) int64 {
	value, ok := o[key]
	if !ok {
		return defaultval
	}
	n, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultval
	}
	return n
}

// Uint returns the value for the given key as uint.
func (o OptionSet) Uint(key string, defaultval uint) uint {
	return uint(o.Uint64(key, uint64(defaultval)))
}

// Uint8 returns the value for the given key as uint8.
func (o OptionSet) Uint8(key string, defaultval int8) uint8 {
	return uint8(o.Uint64(key, uint64(defaultval)))
}

// Uint16 returns the value for the given key as uint16.
func (o OptionSet) Uint16(key string, defaultval int16) uint16 {
	return uint16(o.Uint64(key, uint64(defaultval)))
}

// Uint32 returns the value for the given key as uint32.
func (o OptionSet) Uint32(key string, defaultval uint32) uint32 {
	return uint32(o.Uint64(key, uint64(defaultval)))
}

// Uint64 returns the value for the given key as uint64.
func (o OptionSet) Uint64(key string, defaultval uint64) uint64 {
	value, ok := o[key]
	if !ok {
		return defaultval
	}
	n, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return defaultval
	}
	return n
}

// Float32 returns the value for the given key as float32
func (o OptionSet) Float32(key string, defaultval float32) float32 {
	return float32(o.Float64(key, float64(defaultval)))
}

// Float64 returns the value for the given key as float64
func (o OptionSet) Float64(key string, defaultval float64) float64 {
	value, ok := o[key]
	if !ok {
		return defaultval
	}
	n, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultval
	}
	return n
}

// String returns the value for the given key as string
func (o OptionSet) String(key string, defaultval string) string {
	value, ok := o[key]
	if !ok {
		return defaultval
	}
	return value
}

// Bool returns the value for the given key as bool
func (o OptionSet) Bool(key string, defaultval bool) bool {
	value, ok := o[key]
	if !ok {
		return defaultval
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		return defaultval
	}
	return b
}

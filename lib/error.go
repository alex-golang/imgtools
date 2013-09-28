// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package lib

import "fmt"

// Check panics if the given error is non-nil.
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// With is a convenience wrapper which executes a function
// and ensures that any panics are caught and returned as
// regular errors.
//
// This simply wraps the defer/recover mechanic.
// Any errors which occur in the executed code are expected
// to panic. This makes code easier to write, as it does
// not have to check errors everywhere.
func With(f func()) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("%v", x)
		}
	}()

	f()
	return
}

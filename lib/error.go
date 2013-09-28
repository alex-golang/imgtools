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
// to throw a panic. This makes code easier to follow, as it
// can ommit the error handling code everywhere.
//
// For instance, the following functions (f1, f2, f3) may
// perform any operation that throws a panic when an error occurs.
// The panic is turned into an error, returned by `lib.With`.
//
//    err := lib.With(func() {
//        f1()
//        f2()
//        f3()
//    })
//
func With(f func()) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("%v", x)
		}
	}()

	f()
	return
}

// Package must provides utility functions for handling errors with panics.
package must

// Do takes a value of any type and an error. If the error is not nil, it panics with the given error.
// Otherwise, it returns the value. This function is useful for handling errors that are not expected to occur
// and can be safely ignored or handled with a panic.
//
// Example:
//
//	result := must.Do(someFunction())
//	// If someFunction() returns an error, the program will panic with that error.
//	// Otherwise, the result will be assigned the value returned by someFunction().
func Do[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

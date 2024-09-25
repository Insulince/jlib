package jmust

import (
	"io"
	"reflect"
)

// Must is a generic function that simplifies handling functions that return an error as their last return value.
// If the function returns a non-nil error, Must will panic. Otherwise, it returns the non-error values as a []T.
//
// Example 1: Function with no return value except an error
// For a function in package pkg:
//
//	func Fn() error
//
// Normally, you'd handle the error like this:
//
//	if err := pkg.Fn(); err != nil {
//	    return err
//	}
//
// With Must, you can reduce it to:
//
//	jmust.Must[any](pkg.Fn)
//
// Example 2: Function with a return value and an error
// For a function:
//
//	func Fn(n int) (bool, error)
//
// Normally, you'd handle it like this:
//
//	b, err := pkg.Fn(n)
//	if err != nil {
//	    return err
//	}
//
// Using Must:
//
//	b := jmust.Must[bool](pkg.Fn, n)[0]
//
// Example 3: Function with multiple return values and an error
// For a function:
//
//	func Fn(s string, n int) (bool, float64, error)
//
// Using Must:
//
//	results := jmust.Must[any](pkg.Fn, "test", 42)
//	b := results[0].(bool)
//	f := results[1].(float64)
//
// Must panics if:
// - fn's signature does not return an error as its last return value.
// - fn results in an error when invoked.
//
// A quick aside about T: T represents the type of the return values. There are 4 cases to consider:
//  1. No return values in fn -- In this case T can be anything, since there won't be a returned value, but typically using
//     any will suffice.
//  2. A single return value in fn -- T should be the type of the return value from fn.
//  3. Multiple return values in fn of the same type -- T should be the type of those return values from fn.
//  4. Multiple return values in fn of different types -- T must be any, and unfortunately you will have to use interface
//     conversions to get the results as their rich types, as shown in the third example above.
func Must[T any](fn interface{}, args ...any) []T {
	// Get the reflect.Value of the function
	fnValue := reflect.ValueOf(fn)

	// Ensure fn is a function
	if fnValue.Kind() != reflect.Func {
		panic("Must: fn is not a function")
	}

	// Get the type of the function
	fnType := fnValue.Type()

	// Ensure the function has at least one return value
	if fnType.NumOut() == 0 {
		panic("Must: function must return at least one value")
	}

	// Check if the last return type is an error
	if !fnType.Out(fnType.NumOut() - 1).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		panic("Must: function must return an error as the last return value")
	}

	// Prepare the arguments for the function call
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	// Call the function
	out := fnValue.Call(in)

	// Get the error from the last return value
	err := out[len(out)-1].Interface()
	if err != nil {
		panic(err)
	}

	// Return all results except the error
	returns := make([]T, len(out)-1)
	for i := 0; i < len(out)-1; i++ {
		returns[i] = out[i].Interface().(T)
	}
	return returns
}

// MustClose closes an io.Closer, panicking if an error occurs.
func MustClose(c io.Closer) {
	Must[any](c.Close)
}

// MustWrite writes to an io.Writer, panicking if an error occurs. It returns the number of bytes written.
func MustWrite(w io.Writer, p []byte) int {
	return Must[int](w.Write, p)[0]
}

// MustRead reads from an io.Reader, panicking if an error occurs. It returns the number of bytes read.
func MustRead(r io.Reader, p []byte) int {
	return Must[int](r.Read, p)[0]
}

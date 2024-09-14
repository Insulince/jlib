package jmust

import (
	"reflect"
)

// Must takes a function that returns an error as the last return value and its arguments,
// calls the function, and panics if the function returns an error.
func Must[Args any, Ret any](fn interface{}, args ...Args) []Ret {
	// Get the reflect.Value of the function
	fnValue := reflect.ValueOf(fn)

	// Ensure fn is a function
	if fnValue.Kind() != reflect.Func {
		panic("Must: fn is not a function")
	}

	// Prepare the arguments for the function call
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	// Call the function
	out := fnValue.Call(in)

	// Check if the last return value is an error
	if len(out) == 0 || out[len(out)-1].Kind() != reflect.Interface || !out[len(out)-1].Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		panic("Must: function must return an error as the last return value")
	}

	// Get the error from the last return value
	err := out[len(out)-1].Interface()
	if err != nil {
		panic(err)
	}

	// Return all results except the error
	returns := make([]Ret, len(out)-1)
	for i := 0; i < len(out)-1; i++ {
		returns[i] = out[i].Interface().(Ret)
	}
	return returns
}

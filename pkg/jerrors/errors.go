package jerrors

import (
	"github.com/pkg/errors"
)

var (
	// SomeError represents an arbitrary error intended for use within test
	// files when an error is needed but the content of the error does not
	// matter.
	//
	// Errors are not immutable, but this value should never be overwritten or
	// changed in the slightest.
	SomeError = errors.New("some-error")
)

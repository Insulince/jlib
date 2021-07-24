// Package timer provides a simple Timer interface and implements it via the
// timer struct. The critical idea behind this package is the simple line:
// 		defer Start().Stop(&dur)
// This line captures this intended simplicity of this package, a single line to
// time an arbitrary amount of code. The idea being, the Start() function is
// called immediately upon receipt of this line by the runtime, but the Stop()
// function gets deferred until the end of the surrounding function. And then
// the *time.Duration passed into Stop() is populated with the elapsed duration
// via a named return variable.
//
// You can see all of these ideas come together in the Time function.
package timer

import (
	"time"

	"github.com/pkg/errors"
)

type (
	Timer interface {
		Stop(*time.Duration)
	}

	timer struct {
		start time.Time
	}
)

var (
	_ Timer = (*timer)(nil)
)

// Start creates a new Timer, records when it was started, and returns it.
func Start() Timer {
	return &timer{start: time.Now()}
}

// Stop sets out to the amount of time passed since t.start.
func (t *timer) Stop(out *time.Duration) {
	*out = time.Since(t.start)
}

// Time records the amount of time spent executing fn. If fn errors, Time will
// also error.
//
// Though fn is defined as a func() error, it can execute any arbitrary code,
// simply wrap the desired code in a func() error before providing it here. This
// approach is the same as t he approach taken by the errgroup.Group.Go function
// which also allows arbitrary code to be ran within a wrapper-like function.
func Time(fn func() error) (dur time.Duration, _ error) {
	defer Start().Stop(&dur)
	if err := fn(); err != nil {
		return 0, errors.Wrap(err, "timing function")
	}
	return dur, nil
}

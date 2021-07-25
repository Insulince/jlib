package jtimer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Insulince/jlib/pkg/jerrors"
)

const (
	workDur = 100 * time.Millisecond
	delta   = float64(workDur / 5) // Results should be +/- workDur/5
)

func Test_Start(t *testing.T) {
	t.Parallel()

	t.Run("returned Timer should be a non-nil instance of type *timer", func(t *testing.T) {
		t.Parallel()

		tm := Start()

		require.NotNil(t, tm)
		require.IsType(t, &timer{}, tm)
	})
	t.Run("returned timer should have current time stored as start time", func(t *testing.T) {
		t.Parallel()

		tm := Start()

		require.Equal(t, time.Now(), tm.(*timer).start)
	})
}

func Test_timer_Stop(t *testing.T) {
	t.Parallel()

	t.Run("when Timer is nil", func(t *testing.T) {
		t.Parallel()

		t.Run("when out is not nil, should not touch out", func(t *testing.T) {
			t.Parallel()
			var tm *timer = nil

			time.Sleep(workDur)

			var out time.Duration
			tm.Stop(&out)

			require.Equal(t, time.Duration(0), out)
		})
		t.Run("when out is also nil, should not touch out", func(t *testing.T) {
			t.Parallel()
			var tm *timer = nil

			time.Sleep(workDur)

			tm.Stop(nil)

			// This tests by virtue of not panicking.
		})
	})
	t.Run("when Timer is not nil", func(t *testing.T) {
		t.Parallel()

		t.Run("when out is not nil, should set out to the difference in time between timer.Start and now within delta", func(t *testing.T) {
			t.Parallel()
			tm := Start()

			time.Sleep(workDur)

			var out time.Duration
			tm.Stop(&out)

			assert.InDelta(t, workDur, out, delta)
		})
		t.Run("when out is nil, should not touch out", func(t *testing.T) {
			t.Parallel()
			tm := Start()

			time.Sleep(workDur)

			tm.Stop(nil)

			// This tests by virtue of not panicking.
		})
	})
}

func Test_Time(t *testing.T) {
	t.Parallel()

	t.Run("when fn is nil, should return 0 for duration and should return ErrNilFn", func(t *testing.T) {
		t.Parallel()

		dur, err := Time(nil)

		assert.ErrorIs(t, ErrNilFn, err)
		assert.Equal(t, time.Duration(0), dur)
	})
	t.Run("when fn is not nil", func(t *testing.T) {
		t.Parallel()

		t.Run("when fn errors, should return 0 for duration and should return wrapped error", func(t *testing.T) {
			t.Parallel()
			fn := func() error { return jerrors.SomeError }

			dur, err := Time(fn)

			assert.ErrorIs(t, err, jerrors.SomeError)
			assert.Contains(t, err.Error(), "timing fn")
			assert.Equal(t, time.Duration(0), dur)
		})
		t.Run("when fn does not error, should return duration used within delta and should not return error", func(t *testing.T) {
			t.Parallel()
			fn := func() error { time.Sleep(workDur); return nil }

			dur, err := Time(fn)

			assert.NoError(t, err)
			assert.InDelta(t, workDur, dur, delta)
		})
	})
}

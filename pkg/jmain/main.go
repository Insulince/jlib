package jmain

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"

	"github.com/Insulince/jlib/pkg/jsig"
)

type RunFn func(ctx context.Context) error

func Main(runFn RunFn) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer func(start time.Time) { log.Printf("execution took %v\n", time.Since(start)) }(time.Now())

	done := jsig.Trap()

	log.Println("starting...")

	if err := main(ctx, done, runFn); err != nil {
		program := filepath.Base(os.Args[0])
		return errors.Wrap(err, program)
	}

	log.Println("done")

	return nil
}

func main(ctx context.Context, done <-chan struct{}, runFn RunFn) error {
	select {
	case <-ctx.Done(): // Context was done.
		return errors.Wrap(ctx.Err(), "context done")
	case <-done: // Signal trap indicated we are done.
		return nil
	case err, ok := <-run(ctx, runFn): // An error potentially occurred while running.
		switch {
		case err != nil: // We received an error.
			return errors.Wrap(err, "running")
		case ok: // There is no error and the channel is still open, this is odd behavior.
			return errors.Errorf("nil error sent through error channel")
		default: // There is no error and the channel is closed, this indicates that processing is done.
			return nil
		}
	}
}

func run(ctx context.Context, runFn RunFn) <-chan error {
	errs := make(chan error)

	go func() {
		defer close(errs)

		if err := runFn(ctx); err != nil {
			errs <- err
			return
		}
	}()

	return errs
}

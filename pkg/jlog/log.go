package log

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	ErrNoLoggerFound = errors.New("no logger found")
	ErrNotLogger     = errors.New("logger found of incorrect type")
)

var (
	key struct{}
)

func Standard() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.SetReportCaller(true)
	callerPrettyfier := func(f *runtime.Frame) (function string, file string) {
		return "", fmt.Sprintf(" ./%s:%v", f.File[55:], f.Line)
	}
	textFormatter := logrus.TextFormatter{ForceColors: true, ForceQuote: true, FullTimestamp: true, TimestampFormat: time.RFC3339Nano, PadLevelText: true, QuoteEmptyFields: true, CallerPrettyfier: callerPrettyfier}
	logger.SetFormatter(&textFormatter)
	return logger
}

func RegisterStandard(ctx context.Context) (context.Context, *logrus.Logger) {
	logger := Standard()
	ctx = Register(ctx, logger)
	return ctx, logger
}

func Register(ctx context.Context, logger *logrus.Logger) context.Context {
	ctx = context.WithValue(ctx, key, logger)
	return ctx
}

func FromContext(ctx context.Context) (*logrus.Logger, error) {
	v := ctx.Value(key)
	if v == nil {
		return nil, ErrNoLoggerFound
	}
	logger, ok := v.(*logrus.Logger)
	if !ok {
		return nil, errors.Wrapf(ErrNotLogger, "found %T, should be %T", v, &logrus.Logger{})
	}
	return logger, nil

}

func MustFromContext(ctx context.Context) *logrus.Logger {
	logger, err := FromContext(ctx)
	if err != nil {
		panic(errors.Wrap(err, "must from ctx"))
	}
	return logger
}

func FromContextSafe(ctx context.Context) (context.Context, *logrus.Logger) {
	logger, err := FromContext(ctx)
	if err != nil {
		logger = Standard()
		ctx = Register(ctx, logger)
		return ctx, logger
	}
	return ctx, logger
}

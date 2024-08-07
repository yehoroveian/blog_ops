package log

import (
	"context"

	"go.uber.org/zap"
)

type (
	// Logger provides a general logger interface.
	Logger interface {
		With(args ...Arg) Logger
		Named(name string) Logger
		Apply(ctx context.Context) Logger
		Debug(msg string, args ...Arg)
		Info(msg string, args ...Arg)
		Warn(msg string, args ...Arg)
		Error(msg string, args ...Arg)
		Fatal(msg string, args ...Arg)
		Debugf(format string, args ...interface{})
		Infof(format string, args ...interface{})
		Warnf(format string, args ...interface{})
		Errorf(format string, args ...interface{})
		Fatalf(format string, args ...interface{})
	}
	logger struct {
		log   *zap.SugaredLogger
		apply ContextApplier
		funcs []Middleware
	}
)

// Interface compliance check.
var _ Logger = (*logger)(nil)

func (l *logger) With(args ...Arg) Logger {
	if len(args) == 0 {
		return l
	}
	zapLogger := l.log.With(flatten(args)...)
	return l.child(zapLogger)
}

func (l *logger) Named(name string) Logger {
	if len(name) == 0 {
		return l
	}
	zapLogger := l.log.Named(name)
	return l.child(zapLogger)
}

func (l *logger) Apply(ctx context.Context) Logger {
	ctxArgs := l.apply(ctx)
	return l.With(ctxArgs...)
}

func (l *logger) Debug(msg string, args ...Arg) {
	l.log.Debugw(msg, flatten(args)...)
	l.middleware(DebugLevel, msg)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.log.Debugf(format, args...)
	l.middleware(DebugLevel, format, args...)
}

func (l *logger) Info(msg string, args ...Arg) {
	l.log.Infow(msg, flatten(args)...)
	l.middleware(InfoLevel, msg)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
	l.middleware(InfoLevel, format, args...)
}

func (l *logger) Warn(msg string, args ...Arg) {
	l.log.Warnw(msg, flatten(args)...)
	l.middleware(WarnLevel, msg)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
	l.middleware(WarnLevel, format, args...)
}

func (l *logger) Error(msg string, args ...Arg) {
	l.log.Errorw(msg, flatten(args)...)
	l.middleware(ErrorLevel, msg)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
	l.middleware(ErrorLevel, format, args...)
}

func (l *logger) Fatal(msg string, args ...Arg) {
	l.log.Fatalw(msg, flatten(args)...)
	l.middleware(FatalLevel, msg)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.log.Fatalf(format, args...)
	l.middleware(FatalLevel, format, args...)
}

func (l *logger) middleware(level Level, msg string, args ...interface{}) {
	for _, handle := range l.funcs {
		handle(level, msg, args...)
	}
}

func (l *logger) child(log *zap.SugaredLogger) Logger {
	return &logger{
		log:   log,
		apply: l.apply,
		funcs: l.funcs,
	}
}

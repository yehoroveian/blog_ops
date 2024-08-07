package log

import (
	"context"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Builder interface {
		WithContextApplier(apply ContextApplier) Builder
		WithMiddleware(mw Middleware) Builder
		WithSync(sync io.Writer, level LevelEnablerFunc) Builder
		WithName(name string) Builder
		WithArgs(args ...Arg) Builder
		Build()
	}

	ContextApplier func(ctx context.Context) []Arg

	Middleware func(level Level, msg string, args ...interface{})

	zapCore struct {
		sync  io.Writer
		level LevelEnablerFunc
	}

	builder struct {
		name  string
		args  []Arg
		cores []zapCore
		funcs []Middleware
		apply ContextApplier
	}
)

var _ Builder = (*builder)(nil)

func NewBuilder() Builder {
	return &builder{
		apply: func(ctx context.Context) []Arg {
			return nil
		},
	}
}

func (b *builder) WithContextApplier(apply ContextApplier) Builder {
	b.apply = apply
	return b
}

func (b *builder) WithMiddleware(mw Middleware) Builder {
	b.funcs = append(b.funcs, mw)
	return b
}

func (b *builder) WithSync(sync io.Writer, level LevelEnablerFunc) Builder {
	b.cores = append(b.cores, zapCore{
		sync:  sync,
		level: level,
	})
	return b
}

func (b *builder) WithArgs(args ...Arg) Builder {
	b.args = append(b.args, args...)
	return b
}

func (b *builder) WithName(name string) Builder {
	b.name = name
	return b
}

func (b *builder) Build() {
	// If no cores set, use default ones.
	if len(b.cores) == 0 {
		b.cores = append(b.cores,
			zapCore{sync: Lock(os.Stderr), level: HighPriorityLevels},
			zapCore{sync: Lock(os.Stdout), level: LowPriorityLevels})
	}
	cores := make([]zapcore.Core, len(b.cores))
	for i := range b.cores {
		i := i
		cores[i] = zapcore.NewCore(
			zapcore.NewJSONEncoder(zapcore.EncoderConfig{
				TimeKey:        "timestamp",
				LevelKey:       "severity",
				NameKey:        "logger",
				CallerKey:      "caller",
				MessageKey:     "message",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.CapitalLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			}),
			zapcore.AddSync(b.cores[i].sync),
			zap.LevelEnablerFunc(func(l zapcore.Level) bool {
				return b.cores[i].level(parseZapLevel(l))
			}),
		)
	}

	coreTee := zapcore.NewTee(cores...)
	zapLogger := zap.New(coreTee, zap.AddCaller(), zap.AddCallerSkip(2)).Sugar()

	if len(b.args) > 0 {
		zapLogger = zapLogger.With(flatten(b.args)...)
	}
	if len(b.name) > 0 {
		zapLogger = zapLogger.Named(b.name)
	}

	SetLogger(&logger{
		log:   zapLogger,
		funcs: b.funcs,
		apply: b.apply,
	})
}

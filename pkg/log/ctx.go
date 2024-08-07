package log

import (
	"context"
)

type contextKey struct{}

func FromContext(ctx context.Context) Logger {
	value, ok := ctx.Value(contextKey{}).(Logger)
	if !ok || value == nil {
		// It is border case, probably impossible if you are using platform packages for grpc server, message subscribers
		// etc. where logger always propagated in context.
		return global
	}

	return value
}

func ToContext(ctx context.Context, logger Logger) context.Context {
	ctxWithLogger := context.WithValue(ctx, contextKey{}, logger)

	return ctxWithLogger
}

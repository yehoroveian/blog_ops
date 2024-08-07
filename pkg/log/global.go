package log

import (
	"context"
	"sync"
)

// global logger instance.
var global *logger

// Package logger mutex.
var mx sync.Mutex

func init() {
	// Init default logger.
	NewBuilder().Build()
}

// SetLogger replaces current global logger with given one.
func SetLogger(log Logger) {
	mx.Lock()
	defer mx.Unlock()
	if newGlobal, ok := log.(*logger); ok {
		global = newGlobal
	}
}

// With ...
func With(args ...Arg) Logger {
	if len(args) == 0 {
		return global
	}
	zapLogger := global.log.With(flatten(args)...)
	return global.child(zapLogger)
}

// Named ...
func Named(name string) Logger {
	if len(name) == 0 {
		return global
	}
	zapLogger := global.log.Named(name)
	return global.child(zapLogger)
}

// Apply ...
func Apply(ctx context.Context) Logger {
	ctxArgs := global.apply(ctx)
	return global.With(ctxArgs...)
}

// Debug ...
func Debug(msg string, args ...Arg) {
	global.Debug(msg, args...)
}

// DebugCtx ...
func DebugCtx(ctx context.Context, msg string, args ...Arg) {
	FromContext(ctx).Debug(msg, args...)
}

// Info ...
func Info(msg string, args ...Arg) {
	global.Info(msg, args...)
}

// InfoCtx ...
func InfoCtx(ctx context.Context, msg string, args ...Arg) {
	FromContext(ctx).Info(msg, args...)
}

// Warn ...
func Warn(msg string, args ...Arg) {
	global.Warn(msg, args...)
}

// WarnCtx ...
func WarnCtx(ctx context.Context, msg string, args ...Arg) {
	FromContext(ctx).Warn(msg, args...)
}

// Error ...
func Error(msg string, args ...Arg) {
	global.Error(msg, args...)
}

// ErrorCtx ...
func ErrorCtx(ctx context.Context, msg string, args ...Arg) {
	FromContext(ctx).Error(msg, args...)
}

// Fatal ...
func Fatal(msg string, args ...Arg) {
	global.Fatal(msg, args...)
}

// FatalCtx ...
func FatalCtx(ctx context.Context, msg string, args ...Arg) {
	FromContext(ctx).Fatal(msg, args...)
}

// Debugf ...
func Debugf(format string, args ...interface{}) {
	global.Debugf(format, args...)
}

// DebugfCtx ...
func DebugfCtx(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Debugf(format, args...)
}

// Infof ...
func Infof(format string, args ...interface{}) {
	global.Infof(format, args...)
}

// InfofCtx ...
func InfofCtx(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Infof(format, args...)
}

// Warnf ...
func Warnf(format string, args ...interface{}) {
	global.Warnf(format, args...)
}

// WarnfCtx ...
func WarnfCtx(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Warnf(format, args...)
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	global.Errorf(format, args...)
}

// ErrorfCtx ...
func ErrorfCtx(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Errorf(format, args...)
}

// Fatalf ...
func Fatalf(format string, args ...interface{}) {
	global.Fatalf(format, args...)
}

// FatalfCtx ...
func FatalfCtx(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Fatalf(format, args...)
}

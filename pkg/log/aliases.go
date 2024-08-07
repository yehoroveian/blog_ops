package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Lock(ws WriteSyncer) WriteSyncer {
	return zapcore.Lock(ws)
}

type WriteSyncer = zapcore.WriteSyncer

type Option = zap.Option

type LevelEnablerFunc func(Level) bool

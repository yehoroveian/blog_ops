package log

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

// A Level is a logging priority. Higher levels are more important.
type Level int8

const (
	levelDebug = "debug"
	levelInfo  = "info"
	levelWarn  = "warn"
	levelError = "error"
	levelPanic = "panic"
	levelFatal = "fatal"
)

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in production.
	DebugLevel Level = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

// String returns a lower-case ASCII representation of the log level.
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return levelDebug
	case InfoLevel:
		return levelInfo
	case WarnLevel:
		return levelWarn
	case ErrorLevel:
		return levelError
	case PanicLevel:
		return levelPanic
	case FatalLevel:
		return levelFatal
	default:
		return fmt.Sprintf("Level(%d)", l)
	}
}

func ParseLevel(lvl string) (Level, error) {
	switch lvl {
	case "debug":
		return DebugLevel, nil
	case "info":
		return InfoLevel, nil
	case "warn":
		return WarnLevel, nil
	case "error":
		return ErrorLevel, nil
	case "panic":
		return PanicLevel, nil
	case "fatal":
		return FatalLevel, nil
	default:
		return 0, fmt.Errorf("invalid log valid: %s", lvl)
	}
}

func (l Level) ZapLevel() zapcore.Level {
	switch l {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case PanicLevel:
		return zapcore.PanicLevel
	case FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.InvalidLevel
	}
}

func parseZapLevel(l zapcore.Level) Level {
	switch l {
	case zapcore.DebugLevel:
		return DebugLevel
	case zapcore.InfoLevel:
		return InfoLevel
	case zapcore.WarnLevel:
		return WarnLevel
	case zapcore.ErrorLevel:
		return ErrorLevel
	case zapcore.PanicLevel, zapcore.DPanicLevel:
		return PanicLevel
	case zapcore.FatalLevel:
		return FatalLevel
	default:
		return DebugLevel
	}
}

// Predefined zap.LevelEnablerFunc variables for common use cases.
var (
	LowPriorityLevels = LevelEnablerFunc(func(lvl Level) bool {
		return lvl < ErrorLevel
	})
	HighPriorityLevels = LevelEnablerFunc(func(lvl Level) bool {
		return lvl >= ErrorLevel
	})
)

// LevelEnablerFromLevels makes it easier to use zap.LevelEnablerFunc with logger package.
func LevelEnablerFromLevels(levels ...Level) LevelEnablerFunc {
	return LevelEnablerFunc(func(lvl Level) bool {
		for _, l := range levels {
			if l == lvl {
				return true
			}
		}
		return false
	})
}

package logger

import (
	"context"

	"go.uber.org/zap"
)

var globalLogger = NewDefault("prahari")

// SetGlobalLogger overrides the default global logger.
func SetGlobalLogger(l *Logger) {
	if l != nil {
		globalLogger = l
	}
}

// Info logs structured messages with context.
func Info(ctx context.Context, msg string, fields ...Field) {
	globalLogger.WithContext(ctx).Logger.Info(msg, fields...)
}

// Error logs structured error messages with context.
func Error(ctx context.Context, msg string, fields ...Field) {
	globalLogger.WithContext(ctx).Logger.Error(msg, fields...)
}

// Warn logs structured warning messages with context.
func Warn(ctx context.Context, msg string, fields ...Field) {
	globalLogger.WithContext(ctx).Logger.Warn(msg, fields...)
}

// Debug logs structured debug messages with context.
func Debug(ctx context.Context, msg string, fields ...Field) {
	globalLogger.WithContext(ctx).Logger.Debug(msg, fields...)
}

// Fatal logs structured fatal messages with context.
func Fatal(ctx context.Context, msg string, fields ...Field) {
	globalLogger.WithContext(ctx).Logger.Fatal(msg, fields...)
}

// Err wraps native error messages.
func Err(err error) Field {
	return zap.Error(err)
}

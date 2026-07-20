package logger

import (
	"context"

	"go.uber.org/zap"
)

// Logger is the enterprise wrapper for Zap structured logging.
type Logger struct {
	*zap.Logger
	opts Options
}

// WithContext returns a contextualized child logger extracting correlation and trace IDs.
func (l *Logger) WithContext(ctx context.Context) *Logger {
	if ctx == nil {
		return l
	}
	
	// First extract OpenTelemetry traces if active
	ctx = ExtractTraceContext(ctx)
	
	return &Logger{
		Logger: FromContext(ctx, l.Logger),
		opts:   l.opts,
	}
}

// WithFields adds key-value diagnostic tags to the logger scope.
func (l *Logger) WithFields(fields ...Field) *Logger {
	return &Logger{
		Logger: l.Logger.With(fields...),
		opts:   l.opts,
	}
}

// GetZapLogger exposes the underlying raw Zap instance.
func (l *Logger) GetZapLogger() *zap.Logger {
	return l.Logger
}

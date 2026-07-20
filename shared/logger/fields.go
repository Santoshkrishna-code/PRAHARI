package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Field represents a log structured key-value property.
type Field = zapcore.Field

// String constructs a string field mapping.
func String(key, value string) Field {
	return zap.String(key, value)
}

// Int constructs an integer field mapping.
func Int(key string, value int) Field {
	return zap.Int(key, value)
}

// Duration maps latency metrics.
func Duration(key string, val time.Duration) Field {
	return zap.Duration(key, val)
}

// Any allows dynamic fallback maps serialization.
func Any(key string, val interface{}) Field {
	return zap.Any(key, val)
}

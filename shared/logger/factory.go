package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New instantiates a new contextual Logger using the specified options.
func New(opts Options) (*Logger, error) {
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(opts.Level)); err != nil {
		level = zapcore.InfoLevel // Default fallback
	}

	var encoderConfig zapcore.EncoderConfig
	var encoder zapcore.Encoder

	if opts.Env == "production" {
		// Production JSON output configuration
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.TimeKey = "timestamp"
		encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		// Development human-readable console configuration
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// Output paths setup (standard stdout/stderr)
	output := zapcore.Lock(os.Stdout)
	if len(opts.OutputPaths) > 0 {
		// In production, supports writing to files, custom sockets, or stdout
		if opts.OutputPaths[0] == "stderr" {
			output = zapcore.Lock(os.Stderr)
		}
	}

	core := zapcore.NewCore(
		encoder,
		output,
		level,
	)

	// Build zap logger with Caller frames and stacktrace trigger on Error levels
	zapLogger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)

	// Bind default service name diagnostic tag
	if opts.ServiceName != "" {
		zapLogger = zapLogger.With(zap.String("service", opts.ServiceName))
	}

	return &Logger{
		Logger: zapLogger,
		opts:   opts,
	}, nil
}

// NewDefault returns a standard development console logger.
func NewDefault(serviceName string) *Logger {
	opts := DefaultOptions()
	opts.ServiceName = serviceName
	l, _ := New(opts)
	return l
}

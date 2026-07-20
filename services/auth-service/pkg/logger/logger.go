package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger returns a Zap logger configured for the target environment.
func InitLogger(env, logLevel string) (*zap.Logger, error) {
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		level = zap.DebugLevel // Default fallback
	}

	var encoderConfig zapcore.EncoderConfig
	var encoder zapcore.Encoder

	if env == "production" {
		// Optimized production JSON logging
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.TimeKey = "timestamp"
		encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		// User-friendly development CLI logs
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.Lock(os.Stdout),
		level,
	)

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)), nil
}

package bootstrap

import (
	"context"
	"fmt"

	"prahari/shared/config/sources"
)

type AppConfig struct {
	Port         int    `env:"PORT" envDefault:"8121"`
	Environment  string `env:"ENVIRONMENT" envDefault:"development"`
	DatabaseURL  string `env:"DATABASE_URL"`
	RedisAddr    string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	KafkaBrokers string `env:"KAFKA_BROKERS" envDefault:"localhost:9092"`
	RtspTestUrl  string `env:"RTSP_TEST_URL" envDefault:"rtsp://localhost:8554/live.sdp"`
	ModelPath    string `env:"MODEL_PATH" envDefault:"/models/yolov8.onnx"`
}

func LoadConfig(ctx context.Context) (*AppConfig, error) {
	var cfg AppConfig
	if err := sources.ParseEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	return &cfg, nil
}

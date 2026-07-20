package config

import "time"

type ServerConfig struct {
	HTTPPort     string        `env:"SERVER_HTTP_PORT" envDefault:"8001"`
	GRPCPort     string        `env:"SERVER_GRPC_PORT" envDefault:"9001"`
	ReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT" envDefault:"5s"`
	WriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT" envDefault:"10s"`
	IdleTimeout  time.Duration `env:"SERVER_IDLE_TIMEOUT" envDefault:"120s"`
}

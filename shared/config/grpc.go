package config

// GRPCConfig configures the gRPC service listener port and connection timeouts.
type GRPCConfig struct {
	Port               int `env:"GRPC_PORT" envDefault:"9090" validate:"required,port"`
	MaxConnectionIdle  int `env:"GRPC_MAX_CONN_IDLE_SECS" envDefault:"60" validate:"gt=0"`
	MaxConnectionAge   int `env:"GRPC_MAX_CONN_AGE_SECS" envDefault:"300" validate:"gt=0"`
}

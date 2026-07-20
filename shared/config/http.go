package config

// HTTPConfig defines the HTTP server port and connection properties.
type HTTPConfig struct {
	Port         int `env:"HTTP_PORT" envDefault:"8080" validate:"required,port"`
	ReadTimeout  int `env:"HTTP_READ_TIMEOUT_SECS" envDefault:"5" validate:"gt=0"`
	WriteTimeout int `env:"HTTP_WRITE_TIMEOUT_SECS" envDefault:"10" validate:"gt=0"`
	IdleTimeout  int `env:"HTTP_IDLE_TIMEOUT_SECS" envDefault:"120" validate:"gt=0"`
}

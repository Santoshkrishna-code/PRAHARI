package config

type LoggingConfig struct {
	Level  string `env:"LOG_LEVEL" envDefault:"debug"`
	Format string `env:"LOG_FORMAT" envDefault:"json"` // json or console
}

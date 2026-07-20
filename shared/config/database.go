package config

// DatabaseConfig configures relational database (PostgreSQL/TimescaleDB) connection options.
type DatabaseConfig struct {
	Host                   string `env:"DB_HOST" validate:"required"`
	Port                   int    `env:"DB_PORT" envDefault:"5432" validate:"required,port"`
	User                   string `env:"DB_USER" validate:"required"`
	Password               string `env:"DB_PASSWORD" validate:"required"`
	Name                   string `env:"DB_NAME" validate:"required"`
	SSLMode                string `env:"DB_SSL_MODE" envDefault:"disable" validate:"required,oneof=disable require verify-ca verify-full"`
	MaxOpenConns           int    `env:"DB_MAX_OPEN_CONNS" envDefault:"25" validate:"gte=1"`
	MaxIdleConns           int    `env:"DB_MAX_IDLE_CONNS" envDefault:"10" validate:"gte=1"`
	ConnMaxLifetimeSeconds int    `env:"DB_CONN_MAX_LIFETIME_SECS" envDefault:"1800" validate:"gte=1"`
}

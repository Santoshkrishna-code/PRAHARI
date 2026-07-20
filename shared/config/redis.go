package config

// RedisConfig holds variables for connecting to the ElastiCache/Redis caching cluster.
type RedisConfig struct {
	Address    string `env:"REDIS_ADDRESS" validate:"required"`
	Password   string `env:"REDIS_PASSWORD"`
	DB         int    `env:"REDIS_DB" envDefault:"0" validate:"gte=0"`
	PoolSize   int    `env:"REDIS_POOL_SIZE" envDefault:"50" validate:"gte=1"`
	MaxRetries int    `env:"REDIS_MAX_RETRIES" envDefault:"3" validate:"gte=0"`
}

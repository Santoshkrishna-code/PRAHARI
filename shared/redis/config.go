package redis

// Config holds connection parameters for the Redis cluster.
type Config struct {
	Address    string `json:"address"`
	Password   string `json:"password"`
	DB         int    `json:"db"`
	PoolSize   int    `json:"pool_size"`
	MaxRetries int    `json:"max_retries"`
}

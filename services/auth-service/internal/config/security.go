package config

type SecurityConfig struct {
	AllowedOrigins []string `env:"SECURITY_ALLOWED_ORIGINS" envSeparator:"," envDefault:"*"`
	CSRFKey        string   `env:"SECURITY_CSRF_KEY" envDefault:"prahari_csrf_key_32_bytes_long_!"`
	RateLimitRPS   float64  `env:"SECURITY_RATE_LIMIT_RPS" envDefault:"10"` // Requests Per Second
}

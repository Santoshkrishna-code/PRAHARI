package config

// SecurityConfig holds keys for CORS configurations, rate-limit settings, and cryptographic signing keys.
type SecurityConfig struct {
	AllowedOrigins []string `env:"SECURITY_ALLOWED_ORIGINS" envSeparator:"," validate:"required,gt=0"`
	JWTKeysURL     string   `env:"SECURITY_JWT_KEYS_URL" validate:"omitempty,url"` // JWKS endpoints URL
	JWTIssuer      string   `env:"SECURITY_JWT_ISSUER" validate:"omitempty,url"`
	RateLimitRPS   float64  `env:"SECURITY_RATE_LIMIT_RPS" envDefault:"100.0" validate:"gt=0"`
	CSRFSecret     string   `env:"SECURITY_CSRF_SECRET" validate:"omitempty,min=32"`
}

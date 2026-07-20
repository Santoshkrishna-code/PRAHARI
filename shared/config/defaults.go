package config

// SetDefaults applies backup fallback configurations to the root config instance.
func SetDefaults(cfg *Config) {
	if cfg.Environment == "" {
		cfg.Environment = "development"
	}
	if cfg.ServiceName == "" {
		cfg.ServiceName = "prahari-service"
	}

	// 1. AWS Config Defaults
	if cfg.AWS.Region == "" {
		cfg.AWS.Region = "us-east-1"
	}

	// 2. Database Defaults
	if cfg.Database.Port == 0 {
		cfg.Database.Port = 5432
	}
	if cfg.Database.SSLMode == "" {
		cfg.Database.SSLMode = "disable"
	}
	if cfg.Database.MaxOpenConns == 0 {
		cfg.Database.MaxOpenConns = 25
	}
	if cfg.Database.MaxIdleConns == 0 {
		cfg.Database.MaxIdleConns = 10
	}
	if cfg.Database.ConnMaxLifetimeSeconds == 0 {
		cfg.Database.ConnMaxLifetimeSeconds = 1800
	}

	// 3. Redis Defaults
	if cfg.Redis.PoolSize == 0 {
		cfg.Redis.PoolSize = 50
	}
	if cfg.Redis.MaxRetries == 0 {
		cfg.Redis.MaxRetries = 3
	}

	// 4. HTTP / gRPC Defaults
	if cfg.HTTP.Port == 0 {
		cfg.HTTP.Port = 8080
	}
	if cfg.HTTP.ReadTimeout == 0 {
		cfg.HTTP.ReadTimeout = 5
	}
	if cfg.HTTP.WriteTimeout == 0 {
		cfg.HTTP.WriteTimeout = 10
	}
	if cfg.HTTP.IdleTimeout == 0 {
		cfg.HTTP.IdleTimeout = 120
	}

	if cfg.GRPC.Port == 0 {
		cfg.GRPC.Port = 9090
	}
	if cfg.GRPC.MaxConnectionIdle == 0 {
		cfg.GRPC.MaxConnectionIdle = 60
	}
	if cfg.GRPC.MaxConnectionAge == 0 {
		cfg.GRPC.MaxConnectionAge = 300
	}

	// 5. Telemetry & Security Defaults
	if cfg.Telemetry.PrometheusPort == 0 {
		cfg.Telemetry.PrometheusPort = 9090
	}
	if cfg.Telemetry.SampleRate == 0.0 {
		cfg.Telemetry.SampleRate = 1.0
	}
	if cfg.Security.RateLimitRPS == 0.0 {
		cfg.Security.RateLimitRPS = 100.0
	}
}

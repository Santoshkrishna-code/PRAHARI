package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v9"
)

// Config aggregates all sub-configurations for the Auth Service.
type Config struct {
	Env           string `env:"ENV" envDefault:"development"`
	AWS           AWSConfig
	Cognito       CognitoConfig
	Server        ServerConfig
	Logging       LoggingConfig
	Security      SecurityConfig
	Observability ObservabilityConfig
}

// LoadConfig parses environment variables into the Config struct.
func LoadConfig() (*Config, error) {
	cfg := &Config{}
	
	// caarlos0/env parses environment variables into tagged structures
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config env variables: %w", err)
	}

	// Dynamic validation check
	if cfg.Env == "production" {
		if cfg.Cognito.UserPoolID == "" || cfg.Cognito.ClientID == "" {
			return nil, fmt.Errorf("COGNITO_USER_POOL_ID and COGNITO_CLIENT_ID must be set in production")
		}
		if os.Getenv("AWS_ACCESS_KEY_ID") == "" || os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
			// IAM Roles for Service Accounts (IRSA) will handle this in EKS, but check anyway
			fmt.Fprintln(os.Stderr, "WARNING: AWS credentials variables are empty; relying on IRSA IAM Roles")
		}
	}

	return cfg, nil
}

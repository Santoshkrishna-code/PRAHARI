package config

import (
	"context"
	"fmt"
	"os"

	"prahari/shared/config/sources"
)

// Load parses configuration variables following the standard override precedence hierarchy.
// Allows injecting AWS SSM and Secrets Manager API clients to override settings dynamically.
func Load(ctx context.Context, ssmClient sources.SSMAPI, secretsClient sources.SecretsManagerAPI) (*Config, error) {
	var cfg Config

	// 1. Apply backup Defaults
	SetDefaults(&cfg)

	// 2. Load Local Dotenv (Local Development overrides)
	// Attempts to read ".env" from working directory. Skips silently if file is missing.
	_ = sources.LoadDotEnv(".env")

	// 3. Parse environment variables (Container / Kubernetes overrides)
	if err := sources.ParseEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse environment configurations: %w", err)
	}

	// 4. AWS SSM Parameter Store Overrides (if enabled and client provided)
	if ssmClient != nil {
		// Example: Resolve database host overrides dynamically from SSM Parameters
		if ssmParamName := os.Getenv("SSM_DB_HOST_PARAM"); ssmParamName != "" {
			host, err := sources.LoadFromSSM(ctx, ssmClient, ssmParamName)
			if err == nil {
				cfg.Database.Host = host
			}
		}
	}

	// 5. AWS Secrets Manager Overrides (if enabled and client provided)
	if secretsClient != nil {
		// Example: Resolve database password overrides dynamically from Secrets Manager
		if secretID := os.Getenv("SECRETS_DB_PASSWORD_ID"); secretID != "" {
			password, err := sources.LoadFromSecretsManager(ctx, secretsClient, secretID)
			if err == nil {
				cfg.Database.Password = password
			}
		}
	}

	// 6. Execute structural validations
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return &cfg, nil
}

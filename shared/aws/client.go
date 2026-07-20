package aws

import (
	"context"
	"fmt"

	awsSDK "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// BootstrapConfig loads standard AWS SDK Configurations based on targets.
func BootstrapConfig(ctx context.Context, cfg Config) (awsSDK.Config, error) {
	opts := []func(*config.LoadOptions) error{
		config.WithRegion(cfg.Region),
	}

	// LocalStack testing endpoint integration
	if cfg.EndpointURL != "" {
		opts = append(opts, config.WithEndpointResolverWithOptions(
			awsSDK.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (awsSDK.Endpoint, error) {
				return awsSDK.Endpoint{
					URL:           cfg.EndpointURL,
					SigningRegion: cfg.Region,
				}, nil
			}),
		))
	}

	if cfg.Profile != "" {
		opts = append(opts, config.WithSharedConfigProfile(cfg.Profile))
	}

	awsCfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return awsSDK.Config{}, fmt.Errorf("failed to bootstrap AWS config: %w", err)
	}

	return awsCfg, nil
}

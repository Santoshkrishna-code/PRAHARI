package bootstrap

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	prahariConfig "prahari/services/auth-service/internal/config"
)

// InitCognitoClient builds a Cognito Identity Provider SDK Client.
func InitCognitoClient(ctx context.Context, cfg *prahariConfig.AWSConfig) (*cognitoidentityprovider.Client, error) {
	// Standard AWS V2 SDK loader
	awsOpts := []func(*config.LoadOptions) error{
		config.WithRegion(cfg.Region),
	}

	// For local test mock endpoints injection (like LocalStack)
	if cfg.EndpointURL != "" {
		awsOpts = append(awsOpts, config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           cfg.EndpointURL,
					SigningRegion: cfg.Region,
				}, nil
			}),
		))
	}

	awsCfg, err := config.LoadDefaultConfig(ctx, awsOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to load default AWS config: %w", err)
	}

	return cognitoidentityprovider.NewFromConfig(awsCfg), nil
}

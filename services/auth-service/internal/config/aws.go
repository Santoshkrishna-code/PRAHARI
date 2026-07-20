package config

type AWSConfig struct {
	Region          string `env:"AWS_REGION" envDefault:"us-east-1"`
	EndpointURL     string `env:"AWS_ENDPOINT_URL"` // Optional custom endpoint for mock testing
	AccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
}

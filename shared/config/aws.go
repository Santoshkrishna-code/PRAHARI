package config

// AWSConfig holds connection parameters for AWS service clients.
type AWSConfig struct {
	Region          string `env:"AWS_REGION" envDefault:"us-east-1" validate:"required,aws_region"`
	Profile         string `env:"AWS_PROFILE"`
	RoleARN         string `env:"AWS_ROLE_ARN" validate:"omitempty,aws_arn"`
	AccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
	EndpointURL     string `env:"AWS_ENDPOINT_URL" validate:"omitempty,url"` // Custom endpoint for LocalStack testing
}

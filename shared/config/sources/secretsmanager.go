package sources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// SecretsManagerAPI defines the narrow AWS Secrets Manager Client capability needed to fetch secrets.
type SecretsManagerAPI interface {
	GetSecretValue(ctx context.Context, params *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error)
}

// LoadFromSecretsManager retrieves a single secret string value from AWS Secrets Manager.
func LoadFromSecretsManager(ctx context.Context, client SecretsManagerAPI, secretID string) (string, error) {
	if client == nil {
		return "", fmt.Errorf("secrets manager client is uninitialized")
	}

	input := &secretsmanager.GetSecretValueInput{
		SecretId: &secretID,
	}

	output, err := client.GetSecretValue(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to fetch secret value for %s: %w", secretID, err)
	}

	if output.SecretString == nil {
		return "", fmt.Errorf("secret %s returned empty string", secretID)
	}

	return *output.SecretString, nil
}

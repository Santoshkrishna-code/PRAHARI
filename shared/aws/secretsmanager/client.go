package secretsmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// Client wraps the AWS Secrets Manager SDK client.
type Client struct {
	client *secretsmanager.Client
}

// NewClient constructs a Secrets Manager client wrapper.
func NewClient(client *secretsmanager.Client) *Client {
	return &Client{client: client}
}

// Ping implements aws.HealthChecker checking connectivity by listing secrets.
func (c *Client) Ping(ctx context.Context) error {
	if c.client == nil {
		return fmt.Errorf("secrets manager client is uninitialized")
	}
	_, err := c.client.ListSecrets(ctx, &secretsmanager.ListSecretsInput{MaxResults: &[]int32{1}[0]})
	return err
}

// GetSecret retrieves a raw string secret value.
func (c *Client) GetSecret(ctx context.Context, secretID string) (string, error) {
	if c.client == nil {
		return "", fmt.Errorf("secrets manager client is uninitialized")
	}

	input := &secretsmanager.GetSecretValueInput{
		SecretId: &secretID,
	}

	output, err := c.client.GetSecretValue(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve secret %s: %w", secretID, err)
	}

	if output.SecretString == nil {
		return "", fmt.Errorf("secret %s returned empty payload", secretID)
	}

	return *output.SecretString, nil
}

// GetSecretJSON retrieves a secret JSON string and unmarshals it into the target interface.
func (c *Client) GetSecretJSON(ctx context.Context, secretID string, target interface{}) error {
	secretString, err := c.GetSecret(ctx, secretID)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(secretString), target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON secret payload: %w", err)
	}

	return nil
}

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// Client wraps the AWS SSM SDK client.
type Client struct {
	client *ssm.Client
}

// NewClient constructs an SSM client wrapper.
func NewClient(client *ssm.Client) *Client {
	return &Client{client: client}
}

// Ping implements aws.HealthChecker checking connectivity by listing parameters.
func (c *Client) Ping(ctx context.Context) error {
	if c.client == nil {
		return fmt.Errorf("ssm client is uninitialized")
	}
	_, err := c.client.DescribeParameters(ctx, &ssm.DescribeParametersInput{MaxResults: &[]int32{1}[0]})
	return err
}

// GetParameter retrieves a single string parameter value.
func (c *Client) GetParameter(ctx context.Context, name string) (string, error) {
	if c.client == nil {
		return "", fmt.Errorf("ssm client is uninitialized")
	}

	// Always retrieve decrypted secure strings
	decrypt := true
	input := &ssm.GetParameterInput{
		Name:           &name,
		WithDecryption: &decrypt,
	}

	output, err := c.client.GetParameter(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve SSM parameter %s: %w", name, err)
	}

	if output.Parameter == nil || output.Parameter.Value == nil {
		return "", fmt.Errorf("parameter %s returned empty payload", name)
	}

	return *output.Parameter.Value, nil
}

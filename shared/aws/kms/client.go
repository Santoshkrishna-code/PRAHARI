package kms

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

// Client wraps the AWS KMS SDK client.
type Client struct {
	client *kms.Client
}

// NewClient constructs a KMS client wrapper.
func NewClient(client *kms.Client) *Client {
	return &Client{client: client}
}

// Ping implements aws.HealthChecker checking connectivity by listing keys.
func (c *Client) Ping(ctx context.Context) error {
	if c.client == nil {
		return fmt.Errorf("kms client is uninitialized")
	}
	_, err := c.client.ListKeys(ctx, &kms.ListKeysInput{Limit: &[]int32{1}[0]})
	return err
}

// GenerateDataKey generates a new symmetric data key (AES-256) secured by the master key.
// Returns the plaintext key bytes and the encrypted ciphertext key bytes.
func (c *Client) GenerateDataKey(ctx context.Context, keyID string) ([]byte, []byte, error) {
	if c.client == nil {
		return nil, nil, fmt.Errorf("kms client is uninitialized")
	}

	input := &kms.GenerateDataKeyInput{
		KeyId:   &keyID,
		KeySpec: types.DataKeySpecAes256,
	}

	output, err := c.client.GenerateDataKey(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate KMS data key: %w", err)
	}

	return output.Plaintext, output.CiphertextBlob, nil
}

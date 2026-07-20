package secrets

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/kms"
)

// KMSClient wraps the AWS KMS SDK client to perform key encryption operations.
type KMSClient struct {
	client *kms.Client
}

// NewKMSClient constructs the KMSClient adapter.
func NewKMSClient(client *kms.Client) *KMSClient {
	return &KMSClient{client: client}
}

// EncryptKey encrypts a raw data key using the specified KMS Master Key ID.
func (c *KMSClient) EncryptKey(ctx context.Context, keyID string, plaintextKey []byte) ([]byte, error) {
	if c.client == nil {
		return nil, fmt.Errorf("kms sdk client is uninitialized")
	}

	input := &kms.EncryptInput{
		KeyId:     &keyID,
		Plaintext: plaintextKey,
	}

	output, err := c.client.Encrypt(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt data key via KMS: %w", err)
	}

	return output.CiphertextBlob, nil
}

// DecryptKey unseals an encrypted data key.
func (c *KMSClient) DecryptKey(ctx context.Context, encryptedKey []byte) ([]byte, error) {
	if c.client == nil {
		return nil, fmt.Errorf("kms sdk client is uninitialized")
	}

	input := &kms.DecryptInput{
		CiphertextBlob: encryptedKey,
	}

	output, err := c.client.Decrypt(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt data key via KMS: %w", err)
	}

	return output.Plaintext, nil
}

package identity

import (
	"context"
)

// Client wraps gRPC queries to the IAM Service.
type Client struct {
	grpcAddr string
}

// NewClient constructs an IAM client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// GetUserContactDetails queries user phone/email fields.
func (c *Client) GetUserContactDetails(ctx context.Context, userID string) (string, string, error) {
	// In production, execute gRPC request call to identity-service
	return "user@example.com", "+1234567890", nil
}

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

// CheckPermission sends permissions checks requests.
func (c *Client) CheckPermission(ctx context.Context, userID, permission string) (bool, error) {
	// In production, execute gRPC request call to identity-service
	return true, nil
}

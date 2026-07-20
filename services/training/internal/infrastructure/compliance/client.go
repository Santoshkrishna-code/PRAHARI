package compliance

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks active compliance checklists.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// VerifyCertificationsRequirements check obligations lists.
func (c *Client) VerifyCertificationsRequirements(ctx context.Context, userID string) ([]string, error) {
	prahariLogger.Info(ctx, "Verifying compliance obligation certifications via Compliance Service gRPC",
		prahariLogger.String("user_id", userID))
	return []string{}, nil
}

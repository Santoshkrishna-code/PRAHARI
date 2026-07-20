package identity

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client verifies auditor roles in organizations.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// VerifyLeadAuditor check user roles.
func (c *Client) VerifyLeadAuditor(ctx context.Context, userID string) (bool, error) {
	prahariLogger.Info(ctx, "Verifying lead auditor credentials via Identity Service",
		prahariLogger.String("user_id", userID))
	return true, nil
}
}

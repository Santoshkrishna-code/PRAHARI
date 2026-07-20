package audit

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks active audit findings.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// FetchMandatoryTrainings checks closed findings parameters.
func (c *Client) FetchMandatoryTrainings(ctx context.Context, auditID string) ([]string, error) {
	prahariLogger.Info(ctx, "Querying mandatory training recommendations from closed Audit findings via Audit Service gRPC",
		prahariLogger.String("audit_id", auditID))
	return []string{}, nil
}

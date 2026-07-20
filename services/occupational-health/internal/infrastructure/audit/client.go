package audit

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Client struct {
	grpcAddr string
}

func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

func (c *Client) FetchHealthFindings(ctx context.Context, auditID string) ([]string, error) {
	prahariLogger.Info(ctx, "Retrieving workforce health finding parameters from Audit Service gRPC",
		prahariLogger.String("audit_id", auditID))
	return []string{}, nil
}

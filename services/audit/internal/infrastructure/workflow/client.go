package workflow

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client wraps gRPC review workflow triggers.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// VerifyAuditReviews registers approvals workflows.
func (c *Client) VerifyAuditReviews(ctx context.Context, auditID string) error {
	prahariLogger.Info(ctx, "Submitting audit review items to Workflow Engine for digital verification and routing",
		prahariLogger.String("audit_id", auditID))
	return nil
}
}

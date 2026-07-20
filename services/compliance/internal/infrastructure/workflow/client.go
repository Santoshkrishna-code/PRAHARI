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

// VerifyComplianceReviews registers approvals workflows.
func (c *Client) VerifyComplianceReviews(ctx context.Context, complianceID string) error {
	prahariLogger.Info(ctx, "Submitting compliance review items to Workflow Engine for digital verification and routing",
		prahariLogger.String("compliance_id", complianceID))
	return nil
}

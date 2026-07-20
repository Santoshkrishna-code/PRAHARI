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

// SubmitForApproval registers risk matrix review workflows.
func (c *Client) SubmitForApproval(ctx context.Context, riskID string) error {
	prahariLogger.Info(ctx, "Submitting risk assessment matrix details to Workflow Engine for digital approval verification",
		prahariLogger.String("risk_id", riskID))
	return nil
}

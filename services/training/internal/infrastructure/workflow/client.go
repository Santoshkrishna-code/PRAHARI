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

// VerifyEnrollments registers approvals workflows.
func (c *Client) VerifyEnrollments(ctx context.Context, trainingID string) error {
	prahariLogger.Info(ctx, "Submitting training enrollment items to Workflow Engine for digital verification and routing",
		prahariLogger.String("training_id", trainingID))
	return nil
}

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

// StartWorkflow registers contractor onboarding.
func (c *Client) StartWorkflow(ctx context.Context, contractorID, workflowType string) error {
	prahariLogger.Info(ctx, "Triggering contractor onboarding review workflow in Workflow Engine",
		prahariLogger.String("contractor_id", contractorID),
		prahariLogger.String("workflow_type", workflowType))
	return nil
}

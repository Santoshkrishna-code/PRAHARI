package workflow

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client wraps gRPC orchestrator workflow endpoints.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// StartWorkflow triggers a permit approval flow.
func (c *Client) StartWorkflow(ctx context.Context, permitID, workflowType string) error {
	prahariLogger.Info(ctx, "Triggering approval workflow in Workflow Engine",
		prahariLogger.String("permit_id", permitID),
		prahariLogger.String("workflow_type", workflowType))
	return nil
}

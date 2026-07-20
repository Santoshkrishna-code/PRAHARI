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

// StartWorkflow registers investigation approvals review workflows.
func (c *Client) StartWorkflow(ctx context.Context, nearmissID, workflowType string) error {
	prahariLogger.Info(ctx, "Triggering near miss investigation review workflow in Workflow Engine",
		prahariLogger.String("near_miss_id", nearmissID),
		prahariLogger.String("workflow_type", workflowType))
	return nil
}

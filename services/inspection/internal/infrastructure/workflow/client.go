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

// StartWorkflow registers manager review cycles.
func (c *Client) StartWorkflow(ctx context.Context, inspectionID, workflowType string) error {
	prahariLogger.Info(ctx, "Triggering inspection manager review workflow in Workflow Engine",
		prahariLogger.String("inspection_id", inspectionID),
		prahariLogger.String("workflow_type", workflowType))
	return nil
}

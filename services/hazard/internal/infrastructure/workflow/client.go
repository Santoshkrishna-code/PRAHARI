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

// StartWorkflow registers hazard mitigation reviews.
func (c *Client) StartWorkflow(ctx context.Context, hazardID, workflowType string) error {
	prahariLogger.Info(ctx, "Triggering hazard mitigation plan approval workflow in Workflow Engine",
		prahariLogger.String("hazard_id", hazardID),
		prahariLogger.String("workflow_type", workflowType))
	return nil
}

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

// StartWorkflow registers cost approvals.
func (c *Client) StartWorkflow(ctx context.Context, maintenanceID, workflowType string) error {
	prahariLogger.Info(ctx, "Triggering maintenance cost approval workflow in Workflow Engine",
		prahariLogger.String("maintenance_id", maintenanceID),
		prahariLogger.String("workflow_type", workflowType))
	return nil
}

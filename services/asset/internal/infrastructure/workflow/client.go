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
func (c *Client) StartWorkflow(ctx context.Context, assetID, workflowType string) error {
	prahariLogger.Info(ctx, "Triggering asset commission approval workflow in Workflow Engine",
		prahariLogger.String("asset_id", assetID),
		prahariLogger.String("workflow_type", workflowType))
	return nil
}

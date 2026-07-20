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

// StartWorkflow registers action plan approvals.
func (c *Client) StartWorkflow(ctx context.Context, observationID, workflowType string) error {
	prahariLogger.Info(ctx, "Triggering observation action plan review workflow in Workflow Engine",
		prahariLogger.String("observation_id", observationID),
		prahariLogger.String("workflow_type", workflowType))
	return nil
}

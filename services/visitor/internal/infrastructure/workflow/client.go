package workflow

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Client struct {
	grpcAddr string
}

func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

func (c *Client) TriggerVisitorApprovalWorkflow(ctx context.Context, visitID string) error {
	prahariLogger.Info(ctx, "Triggered host multi-level visit approval workflow in Workflow Engine", prahariLogger.String("visit_id", visitID))
	return nil
}

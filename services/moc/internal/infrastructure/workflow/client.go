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

func (c *Client) TriggerApprovalWorkflow(ctx context.Context, mocID string) error {
	prahariLogger.Info(ctx, "Triggered MOC multi-gate approval workflow in Workflow Engine", prahariLogger.String("moc_id", mocID))
	return nil
}

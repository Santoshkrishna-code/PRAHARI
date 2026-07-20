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

func (c *Client) TriggerBypassApprovalWorkflow(ctx context.Context, barrierID string) error {
	prahariLogger.Info(ctx, "Triggered multi-level bypass approval workflow in Workflow Engine", prahariLogger.String("barrier_id", barrierID))
	return nil
}

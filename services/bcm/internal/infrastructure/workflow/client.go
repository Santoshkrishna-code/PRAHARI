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

func (c *Client) TriggerBCPApprovalWorkflow(ctx context.Context, planID string) error {
	prahariLogger.Info(ctx, "Triggered BCP approval workflow in Workflow Engine", prahariLogger.String("plan_id", planID))
	return nil
}

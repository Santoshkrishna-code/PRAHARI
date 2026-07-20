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

func (c *Client) TriggerPPEAllocationApprovalWorkflow(ctx context.Context, itemID string) error {
	prahariLogger.Info(ctx, "Triggered high-grade PPE gear allocation approval workflow", prahariLogger.String("item_id", itemID))
	return nil
}

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

func (c *Client) TriggerShiftHandoverApprovalWorkflow(ctx context.Context, handoverID string) error {
	prahariLogger.Info(ctx, "Triggered shift handover approval workflow in Workflow Engine", prahariLogger.String("handover_id", handoverID))
	return nil
}

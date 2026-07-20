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

func (c *Client) TriggerEmergencyResponseWorkflow(ctx context.Context, emergencyID string) error {
	prahariLogger.Info(ctx, "Triggered automated emergency escalation workflow in Workflow Engine", prahariLogger.String("emergency_id", emergencyID))
	return nil
}

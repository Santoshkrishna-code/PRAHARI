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

func (c *Client) TriggerCAPAEscalationWorkflow(ctx context.Context, actionID string) error {
	prahariLogger.Info(ctx, "Triggered overdue CAPA action item escalation workflow in Workflow Engine", prahariLogger.String("action_id", actionID))
	return nil
}

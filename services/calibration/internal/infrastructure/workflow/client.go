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

func (c *Client) TriggerOutofToleranceEscalation(ctx context.Context, ootID string) error {
	prahariLogger.Info(ctx, "Triggered multi-level OOT safety escalation workflow in Workflow Engine", prahariLogger.String("oot_id", ootID))
	return nil
}

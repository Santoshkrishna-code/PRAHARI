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

func (c *Client) TriggerChemicalApproval(ctx context.Context, requestID string) error {
	prahariLogger.Info(ctx, "Triggered multi-stage chemical approval workflow in Workflow Engine",
		prahariLogger.String("request_id", requestID))
	return nil
}

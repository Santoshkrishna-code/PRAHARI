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

func (c *Client) SubmitPermitApproval(ctx context.Context, permitID string) error {
	prahariLogger.Info(ctx, "Submitting environmental operating permit to Workflow Engine for digital signature routing",
		prahariLogger.String("permit_id", permitID))
	return nil
}

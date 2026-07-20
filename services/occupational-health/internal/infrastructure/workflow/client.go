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

func (c *Client) SubmitClearanceApproval(ctx context.Context, clearanceID string) error {
	prahariLogger.Info(ctx, "Submitting medical clearance request to Workflow Engine for digital signoff routing",
		prahariLogger.String("clearance_id", clearanceID))
	return nil
}

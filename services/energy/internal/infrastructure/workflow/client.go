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

func (c *Client) SubmitOptimizationApproval(ctx context.Context, recommID string) error {
	prahariLogger.Info(ctx, "Submitting peak-shaving energy optimization to Workflow Engine for approval routing",
		prahariLogger.String("recommendation_id", recommID))
	return nil
}

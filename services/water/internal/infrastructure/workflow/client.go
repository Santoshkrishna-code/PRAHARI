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

func (c *Client) TriggerWaterReviewWorkflow(ctx context.Context, profileID string) error {
	prahariLogger.Info(ctx, "Triggered water stewardship review workflow", prahariLogger.String("profile_id", profileID))
	return nil
}

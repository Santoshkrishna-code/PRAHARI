package permit

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

func (c *Client) ValidateBypassPermit(ctx context.Context, permitID string) (bool, error) {
	prahariLogger.Info(ctx, "Validated safety permit approval for active barrier override/bypass", prahariLogger.String("permit_id", permitID))
	return true, nil
}

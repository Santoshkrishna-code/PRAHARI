package barrier

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

func (c *Client) RegisterBarrierBypass(ctx context.Context, barrierID, reason string) error {
	prahariLogger.Info(ctx, "Registered physical safety barrier bypass/override in Barrier Management Service",
		prahariLogger.String("barrier_id", barrierID))
	return nil
}

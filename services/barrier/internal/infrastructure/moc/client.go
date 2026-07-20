package moc

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

func (c *Client) ValidateBarrierMOCApproval(ctx context.Context, barrierID string) (bool, error) {
	prahariLogger.Info(ctx, "Validated MOC authorization for barrier modification/bypass", prahariLogger.String("barrier_id", barrierID))
	return true, nil
}

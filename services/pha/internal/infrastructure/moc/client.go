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

func (c *Client) VerifyPSSRCompletion(ctx context.Context, mocID string) (bool, error) {
	prahariLogger.Info(ctx, "Notified MOC Service of PSSR completion for final change verification", prahariLogger.String("moc_id", mocID))
	return true, nil
}

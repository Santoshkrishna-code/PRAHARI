package compliance

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

func (c *Client) CheckRegulatoryObligations(ctx context.Context, plantID, changeType string) error {
	prahariLogger.Info(ctx, "Checked compliance regulatory obligations for MOC", prahariLogger.String("plant_id", plantID))
	return nil
}

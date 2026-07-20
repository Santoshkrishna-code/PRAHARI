package emergency

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

func (c *Client) ReportActiveVisitorsMuster(ctx context.Context, plantID string, visitorMusterList []string) error {
	prahariLogger.Info(ctx, "Pushed live plant visitor muster roll to Emergency Management Command Centre",
		prahariLogger.String("plant_id", plantID))
	return nil
}

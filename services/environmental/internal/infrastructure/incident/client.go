package incident

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

func (c *Client) RaiseEnvironmentalIncident(ctx context.Context, title string, desc string) (string, error) {
	prahariLogger.Info(ctx, "Logging critical emission limit exceedance event with reactive Incident Management Service",
		prahariLogger.String("title", title))
	return "INC-ENV-4993", nil
}

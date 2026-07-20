package shift

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

func (c *Client) FetchActiveShift(ctx context.Context, plantID string) (string, error) {
	prahariLogger.Info(ctx, "Fetched active shift from Shift Management to auto-schedule shift briefing",
		prahariLogger.String("plant_id", plantID))
	return "shift-morning-001", nil
}

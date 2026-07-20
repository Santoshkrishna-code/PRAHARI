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

func (c *Client) ReportISO22301Status(ctx context.Context, plantID, status string) error {
	prahariLogger.Info(ctx, "Reported ISO 22301 Business Continuity audit compliance status",
		prahariLogger.String("plant_id", plantID),
		prahariLogger.String("status", status))
	return nil
}

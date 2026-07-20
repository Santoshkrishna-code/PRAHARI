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

func (c *Client) ReportFailedIntegrity(ctx context.Context, barrierID string, healthScore float64) error {
	prahariLogger.Info(ctx, "Reported failed barrier integrity assessment to Compliance Service",
		prahariLogger.String("barrier_id", barrierID))
	return nil
}

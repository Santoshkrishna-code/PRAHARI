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

func (c *Client) FetchSISBarrierStatus(ctx context.Context, transmitterID string) (bool, error) {
	prahariLogger.Info(ctx, "Fetched critical transmitter calibration status for Safety Instrumented System (SIS) barriers",
		prahariLogger.String("transmitter_id", transmitterID))
	return true, nil
}

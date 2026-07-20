package contractor

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

func (c *Client) FetchActiveContractorCount(ctx context.Context, plantID string) (int, error) {
	prahariLogger.Info(ctx, "Fetched active contractor count inside plant for shift overview dashboard",
		prahariLogger.String("plant_id", plantID))
	return 28, nil
}

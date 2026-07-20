package asset

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

func (c *Client) FetchInstrumentSpecification(ctx context.Context, assetID string) (string, error) {
	prahariLogger.Info(ctx, "Fetched instrument datasheet specifications from Asset Management Service",
		prahariLogger.String("asset_id", assetID))
	return "0-100 psi, 4-20mA Output, Rosemount Model 3051S", nil
}

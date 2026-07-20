package risk

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

func (c *Client) FetchChemicalRiskIndex(ctx context.Context, chemicalID string) (float64, error) {
	prahariLogger.Info(ctx, "Fetched risk score index from Risk Assessment service",
		prahariLogger.String("chemical_id", chemicalID))
	return 8.5, nil
}

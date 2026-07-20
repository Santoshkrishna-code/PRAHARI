package hazard

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

func (c *Client) FetchHazardMitigation(ctx context.Context, hazardID string) (string, error) {
	prahariLogger.Info(ctx, "Fetched hazard mitigation plan details to assign preventive action cards",
		prahariLogger.String("hazard_id", hazardID))
	return "Install secondary guard rail at high-level platform elevation", nil
}

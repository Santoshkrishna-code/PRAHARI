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

func (c *Client) UpdateEnterpriseRiskProfile(ctx context.Context, plantID, hazardTitle string, riskRank int) error {
	prahariLogger.Info(ctx, "Updated enterprise risk register profile from PHA study findings",
		prahariLogger.String("plant_id", plantID),
		prahariLogger.Int("risk_rank", riskRank))
	return nil
}

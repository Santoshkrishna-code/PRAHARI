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

func (c *Client) FetchSurveillanceFrequencies(ctx context.Context, riskID string) (int, error) {
	prahariLogger.Info(ctx, "Querying mandatory occupational health program frequencies via Risk Service gRPC",
		prahariLogger.String("risk_id", riskID))
	return 12, nil // e.g. 12 months periodic checks by default
}

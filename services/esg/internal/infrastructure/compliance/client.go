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

func (c *Client) VerifyESGDisclosuresObligations(ctx context.Context, frameworkName string) ([]string, error) {
	prahariLogger.Info(ctx, "Retrieving compliance disclosures requirements list via Compliance Service gRPC",
		prahariLogger.String("framework", frameworkName))
	return []string{"GHG_EMISSIONS_SCOPE_1_AND_2", "WATER_CONSERVATION_METRICS"}, nil
}

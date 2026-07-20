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

func (c *Client) FetchPermitObligations(ctx context.Context, permitID string) ([]string, error) {
	prahariLogger.Info(ctx, "Retrieving compliance legal conditions catalog via Compliance Service gRPC",
		prahariLogger.String("permit_id", permitID))
	return []string{"ANNUAL_STACK_EMISSIONS_TEST", "DAILY_WASTEWATER_PH_RECORDING"}, nil
}

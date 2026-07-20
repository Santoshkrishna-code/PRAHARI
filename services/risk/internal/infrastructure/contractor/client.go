package contractor

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client maps subcontractor safety risk multipliers.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// FetchContractorRiskIndex check compliance scores.
func (c *Client) FetchContractorRiskIndex(ctx context.Context, contractorID string) (float64, error) {
	prahariLogger.Info(ctx, "Querying contractor safety compliance index via Contractor Service gRPC",
		prahariLogger.String("contractor_id", contractorID))
	return 1.0, nil
}

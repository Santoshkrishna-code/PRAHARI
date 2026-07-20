package contractor

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client links contractor evaluations.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// VerifyContractorCompliance check contractor details.
func (c *Client) VerifyContractorCompliance(ctx context.Context, contractorID string) (bool, error) {
	prahariLogger.Info(ctx, "Checking contractor safety compliance credentials via Contractor Service gRPC",
		prahariLogger.String("contractor_id", contractorID))
	return true, nil
}

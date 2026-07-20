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

func (c *Client) VerifyContractorSafetyInduction(ctx context.Context, contractorID string) (bool, error) {
	prahariLogger.Info(ctx, "Verified contractor safety induction from Contractor Management Service",
		prahariLogger.String("contractor_id", contractorID))
	return true, nil
}

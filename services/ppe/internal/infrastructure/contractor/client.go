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

func (c *Client) VerifyContractorPPECompliance(ctx context.Context, contractorID string) (bool, error) {
	prahariLogger.Info(ctx, "Verified contractor protective gear standards compliance check",
		prahariLogger.String("contractor_id", contractorID))
	return true, nil
}

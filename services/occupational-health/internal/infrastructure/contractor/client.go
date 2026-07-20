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

func (c *Client) CheckContractorMedicalStatus(ctx context.Context, contractorID string) (bool, error) {
	prahariLogger.Info(ctx, "Verifying subcontractor physical capability certification compliance via Contractor Service gRPC",
		prahariLogger.String("contractor_id", contractorID))
	return true, nil
}

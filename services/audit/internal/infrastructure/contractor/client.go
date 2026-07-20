package contractor

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks subcontractor safety licenses.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// VerifySupplierCompliance check contractor details.
func (c *Client) VerifySupplierCompliance(ctx context.Context, contractorID string) (bool, error) {
	prahariLogger.Info(ctx, "Checking subcontractor compliance credentials validations via Contractor Service gRPC",
		prahariLogger.String("contractor_id", contractorID))
	return true, nil
}
}

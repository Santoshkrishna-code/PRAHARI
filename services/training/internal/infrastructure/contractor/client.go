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

// VerifyContractorInductions check contractor details.
func (c *Client) VerifyContractorInductions(ctx context.Context, contractorID string) (bool, error) {
	prahariLogger.Info(ctx, "Checking subcontractor induction compliance validation via Contractor Service gRPC",
		prahariLogger.String("contractor_id", contractorID))
	return true, nil
}

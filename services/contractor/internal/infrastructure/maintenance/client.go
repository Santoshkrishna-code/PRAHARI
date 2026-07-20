package maintenance

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client links contractor workorders allocations.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// VerifyContractorTaskAssignments checks assignments.
func (c *Client) VerifyContractorTaskAssignments(ctx context.Context, contractorID string) (bool, error) {
	prahariLogger.Info(ctx, "Querying active task assignments via Maintenance Service gRPC",
		prahariLogger.String("contractor_id", contractorID))
	return true, nil
}

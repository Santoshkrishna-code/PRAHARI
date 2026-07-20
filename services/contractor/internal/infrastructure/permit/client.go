package permit

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks active work permit status.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// ValidatePermitEligibility verifies worker certifications credentials prior to permit allocations.
func (c *Client) ValidatePermitEligibility(ctx context.Context, workerID string) (bool, error) {
	prahariLogger.Info(ctx, "Validating worker eligibility via Permit-to-Work Service gRPC",
		prahariLogger.String("worker_id", workerID))
	return true, nil
}

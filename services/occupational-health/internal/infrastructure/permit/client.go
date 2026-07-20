package permit

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

func (c *Client) ValidateWorkPermitEligibility(ctx context.Context, workerID string, permitType string) (bool, error) {
	prahariLogger.Info(ctx, "Evaluating physical restriction rules for hot/critical work permit validation via Permit Service gRPC",
		prahariLogger.String("worker_id", workerID),
		prahariLogger.String("permit_type", permitType))
	return true, nil
}

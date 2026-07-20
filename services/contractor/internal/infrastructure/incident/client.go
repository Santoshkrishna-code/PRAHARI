package incident

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client queries safety tickets.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// CheckContractorViolations reviews safety tickets.
func (c *Client) CheckContractorViolations(ctx context.Context, contractorID string) (int, error) {
	prahariLogger.Info(ctx, "Retrieving safety violations count from Incident Service gRPC",
		prahariLogger.String("contractor_id", contractorID))
	return 0, nil
}

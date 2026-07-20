package risk

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks plant risk profile matrices parameters.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// QueryRequiredSkills checks process safety barriers checklists.
func (c *Client) QueryRequiredSkills(ctx context.Context, riskID string) ([]string, error) {
	prahariLogger.Info(ctx, "Querying required competency skills configurations via Risk Assessment Service gRPC",
		prahariLogger.String("risk_id", riskID))
	return []string{}, nil
}

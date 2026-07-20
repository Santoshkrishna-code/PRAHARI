package inspection

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks active inspection checklists.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// QueryOnsiteAudits retrieves audits metadata.
func (c *Client) QueryOnsiteAudits(ctx context.Context, contractorID string) ([]string, error) {
	prahariLogger.Info(ctx, "Retrieving audit checks lists from Inspection Service gRPC",
		prahariLogger.String("contractor_id", contractorID))
	return []string{}, nil
}

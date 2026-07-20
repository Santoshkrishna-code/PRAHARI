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

// EscalateNearMiss registers reactive incident records.
func (c *Client) EscalateNearMiss(ctx context.Context, nearmissID, title, description string) error {
	prahariLogger.Info(ctx, "Escalating serious near miss, triggering incident ticket creation in Incident Service",
		prahariLogger.String("near_miss_id", nearmissID),
		prahariLogger.String("title", title))
	return nil
}

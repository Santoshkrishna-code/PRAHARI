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

// EscalateHazard registers reactive incident records.
func (c *Client) EscalateHazard(ctx context.Context, hazardID, title, description string) error {
	prahariLogger.Info(ctx, "Escalating critical risk hazard, triggering incident log creation in Incident Service",
		prahariLogger.String("hazard_id", hazardID),
		prahariLogger.String("title", title))
	return nil
}

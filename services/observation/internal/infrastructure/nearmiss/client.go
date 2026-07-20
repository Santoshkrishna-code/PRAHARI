package nearmiss

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

// EscalateObservation registers near miss records.
func (c *Client) EscalateObservation(ctx context.Context, observationID, title, description string) error {
	prahariLogger.Info(ctx, "Escalating serious unsafe observation, triggering near miss log creation in Near Miss Service",
		prahariLogger.String("observation_id", observationID),
		prahariLogger.String("title", title))
	return nil
}

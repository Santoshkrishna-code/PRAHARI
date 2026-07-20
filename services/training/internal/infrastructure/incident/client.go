package incident

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks reactive records indices.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// FetchRefresherRequests checks incident compliance outcomes.
func (c *Client) FetchRefresherRequests(ctx context.Context, incidentID string) ([]string, error) {
	prahariLogger.Info(ctx, "Retrieving safety refresher training requests from Incident Service gRPC",
		prahariLogger.String("incident_id", incidentID))
	return []string{}, nil
}

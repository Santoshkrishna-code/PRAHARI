package incident

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

func (c *Client) FetchExposureIncidents(ctx context.Context, incidentID string) ([]string, error) {
	prahariLogger.Info(ctx, "Querying chemical/biological hazard exposure records from Incident Service gRPC",
		prahariLogger.String("incident_id", incidentID))
	return []string{}, nil
}

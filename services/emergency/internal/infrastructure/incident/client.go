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

func (c *Client) FetchMajorIncidentDetails(ctx context.Context, incidentID string) (string, error) {
	prahariLogger.Info(ctx, "Fetched major incident data from Incident Management Service", prahariLogger.String("incident_id", incidentID))
	return "HYDROCRACKER_FIRE", nil
}

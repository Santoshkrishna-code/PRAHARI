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

func (c *Client) FetchMajorIncidentImpact(ctx context.Context, incidentID string) (string, error) {
	prahariLogger.Info(ctx, "Fetched major incident business impact assessment", prahariLogger.String("incident_id", incidentID))
	return "CATASTROPHIC", nil
}

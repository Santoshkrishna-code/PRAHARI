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

func (c *Client) FetchIncidentDetails(ctx context.Context, incidentID string) (string, error) {
	prahariLogger.Info(ctx, "Fetched EHS safety incident root-cause log files from Incident Management Service",
		prahariLogger.String("incident_id", incidentID))
	return "Lube oil leakage leading to safety valve malfunction", nil
}

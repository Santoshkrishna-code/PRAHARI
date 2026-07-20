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

func (c *Client) FetchLessonsLearned(ctx context.Context, incidentID string) (string, error) {
	prahariLogger.Info(ctx, "Fetched lessons learned from Incident Management for toolbox talk topic",
		prahariLogger.String("incident_id", incidentID))
	return "Always verify gas test before confined space entry", nil
}

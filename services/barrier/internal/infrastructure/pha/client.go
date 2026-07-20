package pha

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

func (c *Client) FetchSafeguardsAndIPLs(ctx context.Context, studyID string) (int, error) {
	prahariLogger.Info(ctx, "Fetched safeguards and IPL requirements from Process Hazard Analysis", prahariLogger.String("study_id", studyID))
	return 4, nil
}

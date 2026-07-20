package inspection

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks active checklists.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// FetchFieldCompetencyIndicators checklist tags check.
func (c *Client) FetchFieldCompetencyIndicators(ctx context.Context, assetID string) ([]string, error) {
	prahariLogger.Info(ctx, "Retrieving personnel field competency indicators via Inspection Service gRPC",
		prahariLogger.String("asset_id", assetID))
	return []string{}, nil
}

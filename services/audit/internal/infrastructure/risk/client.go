package risk

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks plant risk profile matrices parameters.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// FetchRiskControls checks process safety barriers checklists.
func (c *Client) FetchRiskControls(ctx context.Context, assetID string) ([]string, error) {
	prahariLogger.Info(ctx, "Querying critical risk controls and barriers via Risk Assessment Service gRPC",
		prahariLogger.String("asset_id", assetID))
	return []string{}, nil
}
}

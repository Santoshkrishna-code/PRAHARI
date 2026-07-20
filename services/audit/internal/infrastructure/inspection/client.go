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

// FetchInspectionFindings checklist tags check.
func (c *Client) FetchInspectionFindings(ctx context.Context, assetID string) ([]string, error) {
	prahariLogger.Info(ctx, "Retrieving inspection checklists results and compliance findings via Inspection Service gRPC",
		prahariLogger.String("asset_id", assetID))
	return []string{}, nil
}
}

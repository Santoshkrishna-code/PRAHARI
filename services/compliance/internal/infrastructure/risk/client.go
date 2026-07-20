package risk

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client maps compliance criteria to critical risk controls.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// FetchCriticalControls queries risk assessments controls.
func (c *Client) FetchCriticalControls(ctx context.Context, riskID string) ([]string, error) {
	prahariLogger.Info(ctx, "Retrieving critical risk controls criteria via Risk Assessment Service gRPC",
		prahariLogger.String("risk_id", riskID))
	return []string{}, nil
}

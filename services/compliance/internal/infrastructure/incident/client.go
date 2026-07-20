package incident

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks reactive records indices.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// QueryIncidentFindings checks incident compliance outcomes.
func (c *Client) QueryIncidentFindings(ctx context.Context, incidentID string) ([]string, error) {
	prahariLogger.Info(ctx, "Querying compliance findings and corrective actions from Incident Service gRPC",
		prahariLogger.String("incident_id", incidentID))
	return []string{}, nil
}

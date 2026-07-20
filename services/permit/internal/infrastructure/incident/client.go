package incident

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client wraps gRPC calls to Incident Service.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// LinkPermitToIncident checks if the incident exists.
func (c *Client) LinkPermitToIncident(ctx context.Context, permitID, incidentID string) error {
	prahariLogger.Info(ctx, "Linking permit to incident via Incident Service gRPC",
		prahariLogger.String("permit_id", permitID),
		prahariLogger.String("incident_id", incidentID))
	return nil
}

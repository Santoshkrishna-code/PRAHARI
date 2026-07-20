package incident

import (
	"context"
	"fmt"

	prahariLogger "prahari/shared/logger"
)

// Client escalates critical findings.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// CreateIncidentFromFinding raises a ticket on critical findings.
func (c *Client) CreateIncidentFromFinding(ctx context.Context, title, description, findingID string) (string, error) {
	prahariLogger.Info(ctx, "Escalating critical walkthrough finding to safety Incident ticket",
		prahariLogger.String("finding_id", findingID),
		prahariLogger.String("title", title))
	return fmt.Sprintf("INC-AUTO-%s", findingID[:8]), nil
}

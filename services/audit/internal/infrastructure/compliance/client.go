package compliance

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks active compliance checklists.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// QueryComplianceCriteria check obligations lists.
func (c *Client) QueryComplianceCriteria(ctx context.Context, departmentID string) ([]string, error) {
	prahariLogger.Info(ctx, "Querying active compliance obligations and criteria via Compliance Service gRPC",
		prahariLogger.String("department_id", departmentID))
	return []string{}, nil
}
}

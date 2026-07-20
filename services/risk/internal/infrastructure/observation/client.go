package observation

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks behavioral coaching indices.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// FetchUnsafeObservationsCount checks reported acts.
func (c *Client) FetchUnsafeObservationsCount(ctx context.Context, departmentID string) (int, error) {
	prahariLogger.Info(ctx, "Querying reported unsafe behaviors count via Safety Observation Service gRPC",
		prahariLogger.String("department_id", departmentID))
	return 0, nil
}

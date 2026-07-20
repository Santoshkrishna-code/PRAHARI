package audit

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

func (c *Client) FetchAuditNonConformities(ctx context.Context, auditID string) ([]string, error) {
	prahariLogger.Info(ctx, "Fetched audit non-conformities list to create CAPA items",
		prahariLogger.String("audit_id", auditID))
	return []string{"NC-01: uncalibrated gas detector in plant-B", "NC-02: incomplete LOTO logs"}, nil
}

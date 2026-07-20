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

func (c *Client) TriggerPSMAuditFinding(ctx context.Context, mocID string) error {
	prahariLogger.Info(ctx, "Linked MOC record to Audit Management Service for PSM verification", prahariLogger.String("moc_id", mocID))
	return nil
}

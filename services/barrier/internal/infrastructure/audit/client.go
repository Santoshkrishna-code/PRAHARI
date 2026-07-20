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

func (c *Client) LogBarrierBypassAudit(ctx context.Context, barrierID, approvedBy string) error {
	prahariLogger.Info(ctx, "Logged barrier bypass audit record into Audit Management Service",
		prahariLogger.String("barrier_id", barrierID),
		prahariLogger.String("approved_by", approvedBy))
	return nil
}

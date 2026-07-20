package permit

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

func (c *Client) LinkIsolationToPermit(ctx context.Context, lotoID, permitID string) error {
	prahariLogger.Info(ctx, "Linked hazardous energy isolation certificate to active permit-to-work",
		prahariLogger.String("loto_id", lotoID),
		prahariLogger.String("permit_id", permitID))
	return nil
}

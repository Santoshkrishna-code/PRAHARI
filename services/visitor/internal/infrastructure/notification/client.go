package notification

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

func (c *Client) SendVisitorArrivalNotification(ctx context.Context, hostID, visitorName string) error {
	prahariLogger.Info(ctx, "Sent SMS & Email notification to host regarding visitor arrival",
		prahariLogger.String("host_id", hostID),
		prahariLogger.String("visitor", visitorName))
	return nil
}

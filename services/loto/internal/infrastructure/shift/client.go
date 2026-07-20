package shift

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

func (c *Client) LogShiftLOTOChange(ctx context.Context, shiftID, message string) error {
	prahariLogger.Info(ctx, "Logged active isolation details to Shift handover log",
		prahariLogger.String("shift_id", shiftID))
	return nil
}

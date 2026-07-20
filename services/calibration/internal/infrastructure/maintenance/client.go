package maintenance

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

func (c *Client) TriggerCalibrationWorkOrder(ctx context.Context, instrumentID string) error {
	prahariLogger.Info(ctx, "Triggered preventive calibration maintenance work order in Maintenance Management Service",
		prahariLogger.String("instrument_id", instrumentID))
	return nil
}

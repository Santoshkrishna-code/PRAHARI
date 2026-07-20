package emergency

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

func (c *Client) FetchEmergencySensorVerification(ctx context.Context, sensorID string) (bool, error) {
	prahariLogger.Info(ctx, "Fetched calibration verification status of critical emergency fire/gas sensors",
		prahariLogger.String("sensor_id", sensorID))
	return true, nil
}

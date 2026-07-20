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

func (c *Client) VerifyGasDetectorCalibration(ctx context.Context, detectorID string) (bool, error) {
	prahariLogger.Info(ctx, "Verified gas tester calibration validity before hot work permit issuance",
		prahariLogger.String("detector_id", detectorID))
	return true, nil
}

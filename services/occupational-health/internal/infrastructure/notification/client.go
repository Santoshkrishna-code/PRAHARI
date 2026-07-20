package notification

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SendVaccinationAlert(ctx context.Context, workerID string, vaccineName string) error {
	prahariLogger.Info(ctx, "Dispatching workforce booster/vaccination expiration warning notification via SMS/Email",
		prahariLogger.String("worker_id", workerID),
		prahariLogger.String("vaccine_name", vaccineName))
	return nil
}

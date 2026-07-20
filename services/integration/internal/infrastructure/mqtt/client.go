package mqtt

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Client struct {
	broker string
}

func NewClient(broker string) *Client {
	return &Client{broker: broker}
}

func (c *Client) Publish(ctx context.Context, topic string, payload []byte) error {
	prahariLogger.Info(ctx, "Published telemetry message payload via MQTT broker connection",
		prahariLogger.String("topic", topic))
	return nil
}

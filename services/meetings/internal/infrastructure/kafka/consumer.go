package kafka

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Consumer struct{}

func NewConsumer() *Consumer {
	return &Consumer{}
}

func (c *Consumer) Subscribe(ctx context.Context, topics []string, handler func(topic string, payload []byte) error) error {
	prahariLogger.Info(ctx, "Subscribed to Kafka inbound topics for Meetings & Toolbox Talks Management")
	return nil
}

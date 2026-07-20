package routing

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type OutboundClient interface {
	Publish(ctx context.Context, topic string, payload []byte) error
}

type Service struct {
	client OutboundClient
}

func NewService(client OutboundClient) *Service {
	return &Service{client: client}
}

func (s *Service) RouteMessage(ctx context.Context, targetTopic string, payload []byte) error {
	prahariLogger.Info(ctx, "Routing integration message", prahariLogger.String("topic", targetTopic))
	return s.client.Publish(ctx, targetTopic, payload)
}

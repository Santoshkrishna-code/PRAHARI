package connectors

import (
	"context"
	"fmt"
	"time"

	"prahari/services/integration/internal/domain/connector"
	"prahari/services/integration/internal/domain/events"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveConnector(ctx context.Context, c *connector.Connector) error
	GetConnectorByID(ctx context.Context, id string) (*connector.Connector, error)
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type Service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, pub EventPublisher) *Service {
	return &Service{repo: repo, publisher: pub}
}

func (s *Service) RegisterConnector(ctx context.Context, c *connector.Connector) error {
	c.ID = fmt.Sprintf("conn-%d", time.Now().UnixNano())
	c.Status = "DISCONNECTED"
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	if err := s.repo.SaveConnector(ctx, c); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventConnectorDisconnected, c)
	return nil
}

func (s *Service) TestConnection(ctx context.Context, connectorID string) (bool, error) {
	c, err := s.repo.GetConnectorByID(ctx, connectorID)
	if err != nil {
		return false, err
	}

	// Mock connection check logic
	c.Status = "CONNECTED"
	c.UpdatedAt = time.Now()
	_ = s.repo.SaveConnector(ctx, c)

	_ = s.publisher.Publish(ctx, events.EventConnectorConnected, c)
	prahariLogger.Info(ctx, "Connector connection test successful",
		prahariLogger.String("connector_id", connectorID),
		prahariLogger.String("type", c.Type))
	return true, nil
}

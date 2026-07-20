package registration

import (
	"context"
	"fmt"
	"time"

	"prahari/services/visitor/internal/domain/events"
	"prahari/services/visitor/internal/domain/visitor"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveVisitor(ctx context.Context, vis *visitor.Visitor) error
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type Service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, pub EventPublisher) *Service {
	return &Service{
		repo:      repo,
		publisher: pub,
	}
}

func (s *Service) RegisterVisitor(ctx context.Context, vis *visitor.Visitor) error {
	vis.ID = fmt.Sprintf("vis-%d", time.Now().UnixNano())
	vis.CreatedAt = time.Now()
	vis.UpdatedAt = time.Now()

	if err := s.repo.SaveVisitor(ctx, vis); err != nil {
		return fmt.Errorf("failed to save visitor profile: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventVisitorRegistered, vis)
	prahariLogger.Info(ctx, "Visitor pre-registered successfully",
		prahariLogger.String("first_name", vis.FirstName),
		prahariLogger.String("last_name", vis.LastName))
	return nil
}

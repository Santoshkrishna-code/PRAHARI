package approval

import (
	"context"
	"time"

	"prahari/services/visitor/internal/domain/events"
	"prahari/services/visitor/internal/domain/status"
	"prahari/services/visitor/internal/domain/visit"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetVisitByID(ctx context.Context, id string) (*visit.Visit, error)
	SaveVisit(ctx context.Context, v *visit.Visit) error
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

func (s *Service) ApproveVisit(ctx context.Context, visitID string) error {
	v, err := s.repo.GetVisitByID(ctx, visitID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(v.Status), status.CodeHostApproval); err != nil {
		return err
	}

	v.Status = string(status.CodeHostApproval)
	v.UpdatedAt = time.Now()

	if err := s.repo.SaveVisit(ctx, v); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventVisitorApproved, v)
	prahariLogger.Info(ctx, "Visit hosted approval granted", prahariLogger.String("visit_id", visitID))
	return nil
}

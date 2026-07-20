package checkout

import (
	"context"
	"fmt"
	"time"

	"prahari/services/visitor/internal/domain/checkout"
	"prahari/services/visitor/internal/domain/events"
	"prahari/services/visitor/internal/domain/status"
	"prahari/services/visitor/internal/domain/visit"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetVisitByID(ctx context.Context, id string) (*visit.Visit, error)
	SaveVisit(ctx context.Context, v *visit.Visit) error
	SaveCheckout(ctx context.Context, rec *checkout.Record) error
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

func (s *Service) CheckOut(ctx context.Context, visitID, operatorID string) error {
	v, err := s.repo.GetVisitByID(ctx, visitID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(v.Status), status.CodeCheckedOut); err != nil {
		return err
	}

	rec := &checkout.Record{
		ID:           fmt.Sprintf("out-%d", time.Now().UnixNano()),
		VisitID:      visitID,
		CheckOutAt:   time.Now(),
		CheckedOutBy: operatorID,
		BadgeReturned: true,
	}

	v.Status = string(status.CodeCheckedOut)
	v.UpdatedAt = time.Now()

	if err := s.repo.SaveCheckout(ctx, rec); err != nil {
		return err
	}
	if err := s.repo.SaveVisit(ctx, v); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventVisitorCheckedOut, rec)
	prahariLogger.Info(ctx, "Visitor physically checked out of plant", prahariLogger.String("visit_id", visitID))
	return nil
}

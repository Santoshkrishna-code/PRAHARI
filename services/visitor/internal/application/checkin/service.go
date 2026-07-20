package checkin

import (
	"context"
	"fmt"
	"time"

	"prahari/services/visitor/internal/domain/checkin"
	"prahari/services/visitor/internal/domain/events"
	"prahari/services/visitor/internal/domain/status"
	"prahari/services/visitor/internal/domain/visit"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetVisitByID(ctx context.Context, id string) (*visit.Visit, error)
	SaveVisit(ctx context.Context, v *visit.Visit) error
	SaveCheckin(ctx context.Context, rec *checkin.Record) error
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

func (s *Service) CheckIn(ctx context.Context, visitID, checkpoint, operatorID string) error {
	v, err := s.repo.GetVisitByID(ctx, visitID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(v.Status), status.CodeCheckedIn); err != nil {
		return err
	}

	rec := &checkin.Record{
		ID:                 fmt.Sprintf("chk-%d", time.Now().UnixNano()),
		VisitID:            visitID,
		SecurityCheckPoint: checkpoint,
		CheckInAt:          time.Now(),
		CheckedInBy:        operatorID,
	}

	v.Status = string(status.CodeCheckedIn)
	v.UpdatedAt = time.Now()

	if err := s.repo.SaveCheckin(ctx, rec); err != nil {
		return err
	}
	if err := s.repo.SaveVisit(ctx, v); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventVisitorCheckedIn, rec)
	prahariLogger.Info(ctx, "Visitor physically checked in on-site", prahariLogger.String("visit_id", visitID))
	return nil
}

package scheduling

import (
	"context"
	"fmt"
	"time"

	"prahari/services/shift/internal/domain/events"
	"prahari/services/shift/internal/domain/shift"
	"prahari/services/shift/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveShift(ctx context.Context, sh *shift.Shift) error
	GetShiftByID(ctx context.Context, id string) (*shift.Shift, error)
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

func (s *Service) CreateShift(ctx context.Context, sh *shift.Shift) error {
	sh.ID = fmt.Sprintf("shf-%d", time.Now().UnixNano())
	sh.Status = string(status.CodeScheduled)
	sh.CreatedAt = time.Now()
	sh.UpdatedAt = time.Now()

	if err := s.repo.SaveShift(ctx, sh); err != nil {
		return fmt.Errorf("failed to save shift: %w", err)
	}

	prahariLogger.Info(ctx, "Shift scheduled", prahariLogger.String("shift_name", sh.ShiftName))
	return nil
}

func (s *Service) StartShift(ctx context.Context, id string) error {
	sh, err := s.repo.GetShiftByID(ctx, id)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(sh.Status), status.CodeShiftStarted); err != nil {
		return err
	}

	sh.Status = string(status.CodeShiftStarted)
	now := time.Now()
	sh.ActualStart = &now
	sh.UpdatedAt = now

	if err := s.repo.SaveShift(ctx, sh); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventShiftStarted, sh)
	prahariLogger.Info(ctx, "Shift started", prahariLogger.String("shift_id", id))
	return nil
}

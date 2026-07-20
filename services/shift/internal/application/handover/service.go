package handover

import (
	"context"
	"fmt"
	"time"

	"prahari/services/shift/internal/domain/events"
	"prahari/services/shift/internal/domain/handover"
	"prahari/services/shift/internal/domain/shift"
	"prahari/services/shift/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveHandover(ctx context.Context, ho *handover.Handover) error
	GetHandoverByID(ctx context.Context, id string) (*handover.Handover, error)
	GetShiftByID(ctx context.Context, id string) (*shift.Shift, error)
	SaveShift(ctx context.Context, sh *shift.Shift) error
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

func (s *Service) InitiateHandover(ctx context.Context, ho *handover.Handover) error {
	ho.ID = fmt.Sprintf("hnd-%d", time.Now().UnixNano())
	ho.Status = "PENDING"
	ho.InitiatedAt = time.Now()

	sh, err := s.repo.GetShiftByID(ctx, ho.OutgoingShiftID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(sh.Status), status.CodeHandoverInitiated); err != nil {
		return err
	}

	sh.Status = string(status.CodeHandoverInitiated)
	sh.HandoverID = ho.ID
	sh.UpdatedAt = time.Now()

	if err := s.repo.SaveHandover(ctx, ho); err != nil {
		return err
	}
	if err := s.repo.SaveShift(ctx, sh); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventHandoverInitiated, ho)
	prahariLogger.Warn(ctx, "Shift handover initiated",
		prahariLogger.String("outgoing_shift", ho.OutgoingShiftID),
		prahariLogger.String("incoming_shift", ho.IncomingShiftID))
	return nil
}

func (s *Service) AcceptHandover(ctx context.Context, handoverID string) error {
	ho, err := s.repo.GetHandoverByID(ctx, handoverID)
	if err != nil {
		return err
	}

	sh, err := s.repo.GetShiftByID(ctx, ho.OutgoingShiftID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(sh.Status), status.CodeHandoverAccepted); err != nil {
		return err
	}

	now := time.Now()
	ho.AcceptedAt = &now
	ho.Status = "ACCEPTED"

	sh.Status = string(status.CodeHandoverAccepted)
	sh.UpdatedAt = now

	if err := s.repo.SaveHandover(ctx, ho); err != nil {
		return err
	}
	if err := s.repo.SaveShift(ctx, sh); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventHandoverAccepted, ho)
	prahariLogger.Info(ctx, "Shift handover accepted by incoming supervisor", prahariLogger.String("handover_id", handoverID))
	return nil
}

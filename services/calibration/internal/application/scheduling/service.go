package scheduling

import (
	"context"
	"fmt"
	"time"

	"prahari/services/calibration/internal/domain/calibrationschedule"
	"prahari/services/calibration/internal/domain/events"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveSchedule(ctx context.Context, sched *calibrationschedule.Schedule) error
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

func (s *Service) ScheduleCalibrationTask(ctx context.Context, sched *calibrationschedule.Schedule) error {
	sched.ID = fmt.Sprintf("scd-%d", time.Now().UnixNano())
	sched.Status = "PENDING"

	if err := s.repo.SaveSchedule(ctx, sched); err != nil {
		return fmt.Errorf("failed to save calibration schedule: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventCalibrationScheduled, sched)
	prahariLogger.Info(ctx, "Calibration schedule created",
		prahariLogger.String("instrument_id", sched.InstrumentID))
	return nil
}

package synchronization

import (
	"context"
	"time"

	"prahari/services/digitaltwin/internal/domain/events"
	"prahari/services/digitaltwin/internal/domain/policy"
	"prahari/services/digitaltwin/internal/domain/state"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveLiveState(ctx context.Context, s *state.LiveState) error
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type Service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, publisher EventPublisher) *Service {
	return &Service{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *Service) SyncTelemetry(ctx context.Context, twinID, equipmentID string, val float64, quality string) error {
	if !policy.ValidateTelemetryQuality(quality) {
		prahariLogger.Warn(ctx, "Skipping telemetry update due to poor signal quality",
			prahariLogger.String("equipment_id", equipmentID),
			prahariLogger.String("quality", quality))
		return nil
	}

	st := &state.LiveState{
		ID:          "state-" + equipmentID,
		TwinID:      twinID,
		EquipmentID: equipmentID,
		Value:       val,
		Quality:     quality,
		Timestamp:   time.Now(),
	}

	if err := s.repo.SaveLiveState(ctx, st); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Synchronized live equipment state in Digital Twin",
		prahariLogger.String("twin_id", twinID),
		prahariLogger.String("equipment_id", equipmentID))

	_ = s.publisher.Publish(ctx, events.EventTwinStateUpdated, st)
	return nil
}

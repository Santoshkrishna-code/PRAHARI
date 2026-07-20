package execution

import (
	"context"
	"fmt"
	"time"

	"prahari/services/calibration/internal/domain/calibration"
	"prahari/services/calibration/internal/domain/events"
	"prahari/services/calibration/internal/domain/measurement"
	"prahari/services/calibration/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetCalibrationByID(ctx context.Context, id string) (*calibration.Record, error)
	SaveCalibration(ctx context.Context, rec *calibration.Record) error
	SaveMeasurement(ctx context.Context, m *measurement.Result) error
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

func (s *Service) StartCalibration(ctx context.Context, rec *calibration.Record) error {
	rec.ID = fmt.Sprintf("cal-%d", time.Now().UnixNano())
	rec.Status = string(status.CodeCalibrationStarted)
	rec.CalibratedAt = time.Now()
	rec.CreatedAt = time.Now()
	rec.UpdatedAt = time.Now()

	if err := s.repo.SaveCalibration(ctx, rec); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventCalibrationStarted, rec)
	prahariLogger.Info(ctx, "Calibration execution started", prahariLogger.String("instrument_id", rec.InstrumentID))
	return nil
}

func (s *Service) RecordMeasurements(ctx context.Context, calID string, m *measurement.Result) error {
	rec, err := s.repo.GetCalibrationByID(ctx, calID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(rec.Status), status.CodeMeasurementRecorded); err != nil {
		return err
	}

	m.ID = fmt.Sprintf("msr-%d", time.Now().UnixNano())
	m.CalibrationID = calID
	m.Timestamp = time.Now()

	rec.Status = string(status.CodeMeasurementRecorded)
	rec.UpdatedAt = time.Now()

	if err := s.repo.SaveMeasurement(ctx, m); err != nil {
		return err
	}
	if err := s.repo.SaveCalibration(ctx, rec); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Calibration measurements recorded", prahariLogger.String("calibration_id", calID))
	return nil
}

func (s *Service) ApproveCalibration(ctx context.Context, calID, supervisorID string) error {
	rec, err := s.repo.GetCalibrationByID(ctx, calID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(rec.Status), status.CodeApproved); err != nil {
		return err
	}

	now := time.Now()
	rec.ApprovedBy = supervisorID
	rec.ApprovedAt = &now
	rec.Status = string(status.CodeApproved)
	rec.Result = "PASS"
	rec.UpdatedAt = now

	if err := s.repo.SaveCalibration(ctx, rec); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventCalibrationCompleted, rec)
	prahariLogger.Info(ctx, "Calibration record approved by supervisor", prahariLogger.String("calibration_id", calID))
	return nil
}

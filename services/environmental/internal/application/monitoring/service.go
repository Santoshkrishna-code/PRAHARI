package monitoring

import (
	"context"
	"time"


	"prahari/services/environmental/internal/domain/airquality"
	"prahari/services/environmental/internal/domain/audittrail"
	"prahari/services/environmental/internal/domain/laboratoryresult"
	"prahari/services/environmental/internal/domain/monitoringprogram"
	"prahari/services/environmental/internal/domain/noise"
	"prahari/services/environmental/internal/domain/sampling"
	"prahari/services/environmental/internal/domain/timeline"
	"prahari/services/environmental/internal/domain/vibration"
	"prahari/services/environmental/internal/domain/waterquality"
)

type Repository interface {
	SaveProgram(ctx context.Context, p *monitoringprogram.MonitoringProgram) error
	SaveSampling(ctx context.Context, s *sampling.Sampling) error
	SaveAirQuality(ctx context.Context, a *airquality.AirQuality) error
	SaveWaterQuality(ctx context.Context, w *waterquality.WaterQuality) error
	SaveNoise(ctx context.Context, n *noise.NoiseMonitoring) error
	SaveVibration(ctx context.Context, v *vibration.VibrationMonitoring) error
	SaveLabResult(ctx context.Context, r *laboratoryresult.LaboratoryResult) error
}

type EventPublisher interface {
	Publish(ctx context.Context, topic string, payload interface{}) error
}

type AuditTrailLogger interface {
	LogAuditTrail(ctx context.Context, entry *audittrail.Entry) error
}

type TimelineLogger interface {
	LogTimeline(ctx context.Context, e *timeline.Event) error
}

type Service struct {
	repo     Repository
	events   EventPublisher
	trail    AuditTrailLogger
	timeline TimelineLogger
}

func NewService(repo Repository, events EventPublisher, trail AuditTrailLogger, timeline TimelineLogger) *Service {
	return &Service{
		repo:     repo,
		events:   events,
		trail:    trail,
		timeline: timeline,
	}
}

func (s *Service) CreateProgram(ctx context.Context, p *monitoringprogram.MonitoringProgram) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	if err := p.Validate(); err != nil {
		return err
	}

	if err := s.repo.SaveProgram(ctx, p); err != nil {
		return err
	}

	_ = s.trail.LogAuditTrail(ctx, audittrail.NewEntry("CREATE", "monitoring_program", p.ID, "SYSTEM", nil, map[string]string{"type": p.ProgramType}))
	return nil
}

func (s *Service) RecordSampling(ctx context.Context, sm *sampling.Sampling) error {
	sm.CreatedAt = time.Now()
	sm.UpdatedAt = time.Now()

	if err := sm.Validate(); err != nil {
		return err
	}

	if err := s.repo.SaveSampling(ctx, sm); err != nil {
		return err
	}

	_ = s.events.Publish(ctx, "sampling.completed", sm)
	_ = s.timeline.LogTimeline(ctx, timeline.NewEvent(sm.ID, "SAMPLING_COMPLETED", "SYSTEM", "Sampling records logged successfully", nil))
	return nil
}

func (s *Service) RecordAirQuality(ctx context.Context, a *airquality.AirQuality) error {
	if err := a.Validate(); err != nil {
		return err
	}

	// Trigger compliance alert if AQI exceeds threshold limits
	if a.AQI > 100 {
		a.LimitExceeded = true
		_ = s.events.Publish(ctx, "compliance.failed", map[string]interface{}{
			"id":     a.ID,
			"source": "AIR_QUALITY_STATION",
			"aqi":    a.AQI,
		})
	}

	return s.repo.SaveAirQuality(ctx, a)
}

func (s *Service) RecordWaterQuality(ctx context.Context, w *waterquality.WaterQuality) error {
	if err := w.Validate(); err != nil {
		return err
	}

	if err := s.repo.SaveWaterQuality(ctx, w); err != nil {
		return err
	}

	return nil
}

func (s *Service) RecordNoise(ctx context.Context, n *noise.NoiseMonitoring) error {
	if err := n.Validate(); err != nil {
		return err
	}

	return s.repo.SaveNoise(ctx, n)
}

func (s *Service) RecordVibration(ctx context.Context, v *vibration.VibrationMonitoring) error {
	if err := v.Validate(); err != nil {
		return err
	}

	return s.repo.SaveVibration(ctx, v)
}

func (s *Service) EvaluateLabResult(ctx context.Context, r *laboratoryresult.LaboratoryResult) error {
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()

	if err := r.Validate(); err != nil {
		return err
	}

	if r.AnalyteValue > r.RegulatoryLimit {
		r.IsAbnormal = true
		_ = s.events.Publish(ctx, "laboratory.completed", r)
		_ = s.events.Publish(ctx, "compliance.failed", map[string]interface{}{
			"id":       r.ID,
			"source":   "LAB_RESULT",
			"analyte":  r.AnalyteName,
			"val":      r.AnalyteValue,
			"limit":    r.RegulatoryLimit,
			"severity": "CRITICAL",
		})
	}

	if err := s.repo.SaveLabResult(ctx, r); err != nil {
		return err
	}

	return nil
}

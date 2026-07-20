package forecasting

import (
	"context"
	"time"

	"prahari/services/energy/internal/domain/audittrail"
	"prahari/services/energy/internal/domain/energyforecast"
	"prahari/services/energy/internal/domain/timeline"
)

type Repository interface {
	SaveForecast(ctx context.Context, f *energyforecast.Forecast) error
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

func (s *Service) PredictDemand(ctx context.Context, f *energyforecast.Forecast) error {
	f.GeneratedAt = time.Now()
	if f.ConfidenceRate == 0.0 {
		f.ConfidenceRate = 95.0
	}

	if err := f.Validate(); err != nil {
		return err
	}

	if err := s.repo.SaveForecast(ctx, f); err != nil {
		return err
	}

	_ = s.events.Publish(ctx, "energy.forecast.generated", f)
	_ = s.trail.LogAuditTrail(ctx, audittrail.NewEntry("CREATE", "energy_forecast", f.ID, "SYSTEM", nil, map[string]string{"period": f.ForecastPeriod}))
	_ = s.timeline.LogTimeline(ctx, timeline.NewEvent(f.ID, "FORECAST_GENERATED", "SYSTEM", "Energy demand predictive forecast compiled", nil))

	return nil
}

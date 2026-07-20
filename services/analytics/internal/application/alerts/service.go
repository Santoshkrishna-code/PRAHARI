package alerts

import (
	"context"
	"fmt"
	"time"

	"prahari/services/analytics/internal/domain/alert"
	"prahari/services/analytics/internal/domain/events"
	"prahari/services/analytics/internal/domain/policy"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveAlertRule(ctx context.Context, rule *alert.Rule) error
	GetAlertRules(ctx context.Context, metricKey string) ([]*alert.Rule, error)
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type Service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, pub EventPublisher) *Service {
	return &Service{repo: repo, publisher: pub}
}

func (s *Service) CheckMetricThreshold(ctx context.Context, plantID, key string, val float64) error {
	rules, err := s.repo.GetAlertRules(ctx, key)
	if err != nil {
		return err
	}

	for _, rule := range rules {
		if policy.IsAlertTriggered(rule, val) {
			_ = s.publisher.Publish(ctx, events.EventKPIThresholdExceeded, rule)
			prahariLogger.Error(ctx, "HSE Metric Alert threshold violated",
				prahariLogger.String("metric", key),
				prahariLogger.Any("value", val),
				prahariLogger.Any("threshold", rule.Threshold))
		}
	}
	return nil
}

func (s *Service) CreateAlertRule(ctx context.Context, rule *alert.Rule) error {
	rule.ID = fmt.Sprintf("al-%d", time.Now().UnixNano())
	rule.UpdatedAt = time.Now()
	return s.repo.SaveAlertRule(ctx, rule)
}

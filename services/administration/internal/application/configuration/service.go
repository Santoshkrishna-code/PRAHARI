package configuration

import (
	"context"
	"fmt"
	"time"

	"prahari/services/administration/internal/domain/configuration"
	"prahari/services/administration/internal/domain/events"
	"prahari/services/administration/internal/domain/featureflag"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveConfiguration(ctx context.Context, param *configuration.Param) error
	GetConfiguration(ctx context.Context, tenantID, key string) (*configuration.Param, error)
	SaveFeatureFlag(ctx context.Context, flag *featureflag.Flag) error
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

func (s *Service) UpdateConfiguration(ctx context.Context, param *configuration.Param) error {
	param.ID = fmt.Sprintf("cfg-%d", time.Now().UnixNano())
	param.UpdatedAt = time.Now()

	if err := s.repo.SaveConfiguration(ctx, param); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventConfigurationUpdated, param)
	prahariLogger.Info(ctx, "Configuration updated", prahariLogger.String("config_key", param.ConfigKey))
	return nil
}

func (s *Service) GetConfiguration(ctx context.Context, tenantID, key string) (*configuration.Param, error) {
	return s.repo.GetConfiguration(ctx, tenantID, key)
}

func (s *Service) SetFeatureFlag(ctx context.Context, flag *featureflag.Flag) error {
	flag.ID = fmt.Sprintf("ff-%d", time.Now().UnixNano())
	flag.UpdatedAt = time.Now()

	if err := s.repo.SaveFeatureFlag(ctx, flag); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventFeatureflagChanged, flag)
	return nil
}

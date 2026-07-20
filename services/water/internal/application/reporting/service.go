package reporting

import (
	"context"

	"prahari/services/water/internal/domain/waterprofile"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveProfile(ctx context.Context, profile *waterprofile.Profile) error
	GetProfileByID(ctx context.Context, id string) (*waterprofile.Profile, error)
	ListProfiles(ctx context.Context, plantID string) ([]*waterprofile.Profile, error)
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

func (s *Service) CreateProfile(ctx context.Context, profile *waterprofile.Profile) error {
	if err := s.repo.SaveProfile(ctx, profile); err != nil {
		return err
	}
	prahariLogger.Info(ctx, "Water profile created", prahariLogger.String("id", profile.ID))
	return nil
}

func (s *Service) GetProfile(ctx context.Context, id string) (*waterprofile.Profile, error) {
	return s.repo.GetProfileByID(ctx, id)
}

func (s *Service) ListProfiles(ctx context.Context, plantID string) ([]*waterprofile.Profile, error) {
	return s.repo.ListProfiles(ctx, plantID)
}

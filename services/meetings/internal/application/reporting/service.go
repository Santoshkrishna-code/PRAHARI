package reporting

import (
	"context"

	"prahari/services/meetings/internal/domain/meeting"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetMeetingByID(ctx context.Context, id string) (*meeting.Meeting, error)
	ListMeetings(ctx context.Context, plantID string) ([]*meeting.Meeting, error)
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetMeeting(ctx context.Context, id string) (*meeting.Meeting, error) {
	return s.repo.GetMeetingByID(ctx, id)
}

func (s *Service) ListMeetings(ctx context.Context, plantID string) ([]*meeting.Meeting, error) {
	return s.repo.ListMeetings(ctx, plantID)
}

func (s *Service) GetExecutiveMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving safety communication dashboard metrics",
		prahariLogger.String("plant_id", plantID))
	return s.repo.GetDashboardMetrics(ctx, plantID)
}

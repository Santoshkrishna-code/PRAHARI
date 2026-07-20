package search

import (
	"context"

	"prahari/services/meetings/internal/domain/meeting"
	"prahari/services/meetings/internal/domain/search"
)

type Repository interface {
	SearchMeetings(ctx context.Context, criteria *search.Criteria) ([]*meeting.Meeting, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*meeting.Meeting, int64, error) {
	return s.repo.SearchMeetings(ctx, criteria)
}

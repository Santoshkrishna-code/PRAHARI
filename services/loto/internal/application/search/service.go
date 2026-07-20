package search

import (
	"context"

	"prahari/services/loto/internal/domain/isolationcertificate"
	"prahari/services/loto/internal/domain/search"
)

type Repository interface {
	SearchCertificates(ctx context.Context, criteria *search.Criteria) ([]*isolationcertificate.Certificate, int64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExecuteSearch(ctx context.Context, criteria *search.Criteria) ([]*isolationcertificate.Certificate, int64, error) {
	return s.repo.SearchCertificates(ctx, criteria)
}

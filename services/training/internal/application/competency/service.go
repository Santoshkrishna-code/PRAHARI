package competency

import (
	"context"

	matrixDomain "prahari/services/training/internal/domain/competencymatrix"
)

// Repository persistence matrices.
type Repository interface {
	Create(ctx context.Context, cm *matrixDomain.CompetencyMatrix) error
}

// Service manages competencies maps.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// DefineCompetency maps role attributes.
func (s *Service) DefineCompetency(ctx context.Context, cm *matrixDomain.CompetencyMatrix) (*matrixDomain.CompetencyMatrix, error) {
	if err := cm.Validate(); err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, cm); err != nil {
		return nil, err
	}
	return cm, nil
}

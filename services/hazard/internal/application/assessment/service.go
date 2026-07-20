package assessment

import (
	"context"

	matrixDomain "prahari/services/hazard/internal/domain/riskmatrix"
)

// Repository persistence matrix definitions.
type Repository interface {
	Create(ctx context.Context, rm *matrixDomain.RiskMatrix) error
	FindByID(ctx context.Context, id string) (*matrixDomain.RiskMatrix, error)
}

// Service manages 5x5 assessments.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CalculateRisk Score logic helper.
func (s *Service) CalculateRisk(ctx context.Context, rm *matrixDomain.RiskMatrix) (int, string, error) {
	if err := rm.Validate(); err != nil {
		return 0, "", err
	}
	return rm.CalculateRiskScore(), string(rm.GetRiskLevel()), nil
}

package storage

import (
	"context"


	"prahari/services/chemical/internal/domain/compatibility"
	"prahari/services/chemical/internal/domain/container"
	"prahari/services/chemical/internal/domain/policy"
	"prahari/services/chemical/internal/domain/storagearea"
)

type Repository interface {
	GetStorageAreaByID(ctx context.Context, id string) (*storagearea.Area, error)
	GetCompatibilityRules(ctx context.Context) ([]*compatibility.Rule, error)
	SaveContainer(ctx context.Context, con *container.Container) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) StoreContainer(ctx context.Context, con *container.Container, areaID string) error {
	area, err := s.repo.GetStorageAreaByID(ctx, areaID)
	if err != nil {
		return err
	}

	// Verify maximum allowable quantity policy
	if err := policy.VerifyMaxAllowableQuantity(area, con.Capacity); err != nil {
		return err
	}

	con.StorageAreaID = areaID
	con.Status = "STORED"

	return s.repo.SaveContainer(ctx, con)
}

func (s *Service) CheckCompatibility(ctx context.Context, classA, classB string) (bool, error) {
	rules, err := s.repo.GetCompatibilityRules(ctx)
	if err != nil {
		return false, err
	}
	return policy.IsCompatible(classA, classB, rules), nil
}

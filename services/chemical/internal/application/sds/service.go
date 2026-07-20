package sds

import (
	"context"
	"fmt"
	"time"

	sdsDomain "prahari/services/chemical/internal/domain/sds"
	"prahari/services/chemical/internal/domain/sdsrevision"
)

type Repository interface {
	SaveSDS(ctx context.Context, s *sdsDomain.SDS) error
	GetSDSByChemicalID(ctx context.Context, chemicalID string) (*sdsDomain.SDS, error)
	SaveSDSRevision(ctx context.Context, rev *sdsrevision.Revision) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) RegisterSDS(ctx context.Context, sd *sdsDomain.SDS) error {
	sd.ID = fmt.Sprintf("sds-%d", time.Now().UnixNano())
	sd.CreatedAt = time.Now()
	sd.UpdatedAt = time.Now()

	return s.repo.SaveSDS(ctx, sd)
}

func (s *Service) AddRevision(ctx context.Context, sdsID string, rev *sdsrevision.Revision) error {
	rev.ID = fmt.Sprintf("rev-%d", time.Now().UnixNano())
	rev.SdsID = sdsID
	rev.RevisedAt = time.Now()

	return s.repo.SaveSDSRevision(ctx, rev)
}

func (s *Service) GetSDS(ctx context.Context, chemicalID string) (*sdsDomain.SDS, error) {
	return s.repo.GetSDSByChemicalID(ctx, chemicalID)
}

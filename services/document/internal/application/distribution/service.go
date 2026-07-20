package distribution

import (
	"context"
	"fmt"
	"time"

	"prahari/services/document/internal/domain/controlledcopy"
	"prahari/services/document/internal/domain/documentdistribution"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveDistribution(ctx context.Context, rec *documentdistribution.Record) error
	SaveControlledCopy(ctx context.Context, copy *controlledcopy.Copy) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) IssueControlledCopy(ctx context.Context, copy *controlledcopy.Copy) error {
	copy.ID = fmt.Sprintf("cc-%d", time.Now().UnixNano())
	copy.CopyNumber = fmt.Sprintf("CC-%d", time.Now().Unix()%1000)
	copy.Status = "ACTIVE"
	copy.IssuedAt = time.Now()

	if err := s.repo.SaveControlledCopy(ctx, copy); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Watermarked controlled copy issued", prahariLogger.String("copy_number", copy.CopyNumber))
	return nil
}

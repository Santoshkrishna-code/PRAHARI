package service

import (
	"context"
	"fmt"
	"time"

	"prahari/templates/service-template/internal/domain"
)

type PermitUseCaseImpl struct {
	repo domain.PermitRepository
}

func NewPermitUseCase(repo domain.PermitRepository) domain.PermitUseCase {
	return &PermitUseCaseImpl{repo: repo}
}

func (s *PermitUseCaseImpl) GetPermit(ctx context.Context, id string) (*domain.Permit, error) {
	if id == "" {
		return nil, fmt.Errorf("permit ID cannot be empty")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *PermitUseCaseImpl) RequestPermit(ctx context.Context, workerID, zoneID string, duration time.Duration) (*domain.Permit, error) {
	if workerID == "" || zoneID == "" {
		return nil, fmt.Errorf("worker ID and zone ID are required")
	}
	
	permit := &domain.Permit{
		ID:        fmt.Sprintf("ptw-%d", time.Now().UnixNano()),
		WorkerID:  workerID,
		ZoneID:    zoneID,
		Status:    "PENDING",
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}
	
	err := s.repo.Create(ctx, permit)
	if err != nil {
		return nil, err
	}
	
	return permit, nil
}

func (s *PermitUseCaseImpl) ApprovePermit(ctx context.Context, id string, supervisorID string) error {
	if id == "" || supervisorID == "" {
		return fmt.Errorf("permit ID and supervisor ID are required")
	}
	
	permit, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	
	if permit.Status != "PENDING" {
		return fmt.Errorf("permit %s cannot be approved because it is in %s status", id, permit.Status)
	}
	
	permit.Status = "APPROVED"
	permit.ApprovedBy = supervisorID
	permit.ExpiresAt = time.Now().Add(8 * time.Hour) // Standard 8hr shift
	
	return s.repo.Update(ctx, permit)
}

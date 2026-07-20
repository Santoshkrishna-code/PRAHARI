package prooftest

import (
	"context"
	"fmt"
	"time"

	"prahari/services/barrier/internal/domain/barrier"
	"prahari/services/barrier/internal/domain/events"
	"prahari/services/barrier/internal/domain/prooftest"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetBarrierByID(ctx context.Context, id string) (*barrier.Barrier, error)
	SaveBarrier(ctx context.Context, b *barrier.Barrier) error
	SaveProofTest(ctx context.Context, pt *prooftest.Test) error
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type MaintenanceClient interface {
	CreateWorkOrder(ctx context.Context, plantID, title, description string) (string, error)
}

type Service struct {
	repo        Repository
	publisher   EventPublisher
	maintClient MaintenanceClient
}

func NewService(repo Repository, pub EventPublisher, maintClient MaintenanceClient) *Service {
	return &Service{
		repo:        repo,
		publisher:   pub,
		maintClient: maintClient,
	}
}

func (s *Service) RecordProofTest(ctx context.Context, pt *prooftest.Test) error {
	b, err := s.repo.GetBarrierByID(ctx, pt.BarrierID)
	if err != nil {
		return err
	}

	pt.ID = fmt.Sprintf("pt-%d", time.Now().UnixNano())
	pt.TestNumber = fmt.Sprintf("PT-%d", time.Now().Unix()%100000)
	pt.ExecutedAt = time.Now()

	now := time.Now()
	b.LastProofTestedAt = &now
	nextDue := now.AddDate(1, 0, 0) // Default 1 year proof test interval
	b.NextProofTestDue = &nextDue

	if !pt.Passed && s.maintClient != nil {
		woID, err := s.maintClient.CreateWorkOrder(ctx, b.PlantID, fmt.Sprintf("Proof Test Repair for Barrier: %s", b.BarrierCode), pt.Notes)
		if err == nil {
			pt.WorkOrderID = woID
		}
	}

	if err := s.repo.SaveProofTest(ctx, pt); err != nil {
		return err
	}

	_ = s.repo.SaveBarrier(ctx, b)
	_ = s.publisher.Publish(ctx, events.EventProofTestCompleted, pt)
	prahariLogger.Info(ctx, "Proof test recorded", prahariLogger.String("test_number", pt.TestNumber), prahariLogger.String("passed", fmt.Sprintf("%t", pt.Passed)))
	return nil
}


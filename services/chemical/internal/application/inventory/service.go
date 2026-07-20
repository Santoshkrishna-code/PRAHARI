package inventory

import (
	"context"
	"fmt"
	"time"

	"prahari/services/chemical/internal/domain/container"
	"prahari/services/chemical/internal/domain/events"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveContainer(ctx context.Context, con *container.Container) error
	GetContainerByID(ctx context.Context, id string) (*container.Container, error)
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type Service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, pub EventPublisher) *Service {
	return &Service{
		repo:      repo,
		publisher: pub,
	}
}

func (s *Service) ReceiveContainer(ctx context.Context, con *container.Container) error {
	con.ID = fmt.Sprintf("con-%d", time.Now().UnixNano())
	con.Status = "RECEIVED"
	con.CreatedAt = time.Now()
	con.UpdatedAt = time.Now()

	if err := s.repo.SaveContainer(ctx, con); err != nil {
		return fmt.Errorf("failed to receive container: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventChemicalReceived, con)
	prahariLogger.Info(ctx, "Chemical container received",
		prahariLogger.String("container_id", con.ID),
		prahariLogger.String("barcode", con.Barcode))
	return nil
}

func (s *Service) IssueContainer(ctx context.Context, containerID, issuedTo string) error {
	con, err := s.repo.GetContainerByID(ctx, containerID)
	if err != nil {
		return err
	}

	con.Status = "ISSUED"
	con.UpdatedAt = time.Now()

	if err := s.repo.SaveContainer(ctx, con); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventChemicalIssued, con)
	prahariLogger.Info(ctx, "Chemical container issued",
		prahariLogger.String("container_id", containerID),
		prahariLogger.String("issued_to", issuedTo))
	return nil
}

func (s *Service) ReturnContainer(ctx context.Context, containerID string) error {
	con, err := s.repo.GetContainerByID(ctx, containerID)
	if err != nil {
		return err
	}

	con.Status = "STORED"
	con.UpdatedAt = time.Now()

	return s.repo.SaveContainer(ctx, con)
}

func (s *Service) TransferContainer(ctx context.Context, containerID, targetAreaID string) error {
	con, err := s.repo.GetContainerByID(ctx, containerID)
	if err != nil {
		return err
	}

	con.StorageAreaID = targetAreaID
	con.Status = "STORED"
	con.UpdatedAt = time.Now()

	if err := s.repo.SaveContainer(ctx, con); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventChemicalTransferred, con)
	return nil
}

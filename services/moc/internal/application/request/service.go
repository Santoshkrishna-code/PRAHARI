package request

import (
	"context"
	"fmt"
	"time"

	"prahari/services/moc/internal/domain/changerequest"
	"prahari/services/moc/internal/domain/events"
	"prahari/services/moc/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveRequest(ctx context.Context, req *changerequest.Request) error
	GetRequestByID(ctx context.Context, id string) (*changerequest.Request, error)
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

func (s *Service) CreateRequest(ctx context.Context, req *changerequest.Request) error {
	req.ID = fmt.Sprintf("moc-%d", time.Now().UnixNano())
	req.MOCNumber = fmt.Sprintf("MOC-%s-%d", req.PlantID, time.Now().Unix()%100000)
	req.Status = string(status.CodeDraft)
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	if err := s.repo.SaveRequest(ctx, req); err != nil {
		return fmt.Errorf("failed to save change request: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventMOCCreated, req)
	prahariLogger.Info(ctx, "MOC request created", prahariLogger.String("moc_number", req.MOCNumber))
	return nil
}

func (s *Service) GetRequest(ctx context.Context, id string) (*changerequest.Request, error) {
	return s.repo.GetRequestByID(ctx, id)
}

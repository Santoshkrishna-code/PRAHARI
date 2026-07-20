package recommendation

import (
	"context"
	"fmt"
	"time"

	"prahari/services/pha/internal/domain/actionitem"
	"prahari/services/pha/internal/domain/events"
	"prahari/services/pha/internal/domain/recommendation"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveRecommendation(ctx context.Context, rec *recommendation.Recommendation) error
	SaveActionItem(ctx context.Context, item *actionitem.ActionItem) error
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

func (s *Service) CreateRecommendation(ctx context.Context, rec *recommendation.Recommendation) error {
	rec.ID = fmt.Sprintf("rec-%d", time.Now().UnixNano())
	rec.RecNumber = fmt.Sprintf("REC-%d", time.Now().Unix()%100000)
	rec.Status = "OPEN"
	rec.CreatedAt = time.Now()
	rec.UpdatedAt = time.Now()

	if err := s.repo.SaveRecommendation(ctx, rec); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventRecommendationCreated, rec)
	prahariLogger.Info(ctx, "PHA recommendation created", prahariLogger.String("rec_number", rec.RecNumber))
	return nil
}

func (s *Service) CreateActionItem(ctx context.Context, item *actionitem.ActionItem) error {
	item.ID = fmt.Sprintf("act-%d", time.Now().UnixNano())
	item.Status = "OPEN"
	item.CreatedAt = time.Now()

	if s.maintClient != nil && item.WorkOrderID == "" {
		woID, err := s.maintClient.CreateWorkOrder(ctx, "P01", fmt.Sprintf("PHA Action: %s", item.ActionTitle), "Process hazard action item implementation")
		if err == nil {
			item.WorkOrderID = woID
		}
	}

	return s.repo.SaveActionItem(ctx, item)
}

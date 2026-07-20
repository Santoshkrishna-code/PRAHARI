package response

import (
	"context"
	"fmt"
	"time"

	"prahari/services/emergency/internal/domain/emergency"
	"prahari/services/emergency/internal/domain/events"
	"prahari/services/emergency/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)


type Repository interface {
	SaveEmergency(ctx context.Context, em *emergency.Emergency) error
	GetEmergencyByID(ctx context.Context, id string) (*emergency.Emergency, error)
	DeployResource(ctx context.Context, resID string, qty int) error
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type NotificationClient interface {
	SendEmergencyAlert(ctx context.Context, plantID, title, severity string) error
}

type Service struct {
	repo       Repository
	publisher  EventPublisher
	notifClient NotificationClient
}

func NewService(repo Repository, pub EventPublisher, notifClient NotificationClient) *Service {
	return &Service{
		repo:        repo,
		publisher:   pub,
		notifClient: notifClient,
	}
}

func (s *Service) DeclareEmergency(ctx context.Context, em *emergency.Emergency) error {
	em.ID = fmt.Sprintf("emg-%d", time.Now().UnixNano())
	em.EmergencyNumber = fmt.Sprintf("EMG-%s-%d", em.PlantID, time.Now().Unix()%100000)
	em.Status = string(status.CodeDeclared)
	em.DeclaredAt = time.Now()
	em.CreatedAt = time.Now()
	em.UpdatedAt = time.Now()

	if err := s.repo.SaveEmergency(ctx, em); err != nil {
		return err
	}

	if s.notifClient != nil {
		_ = s.notifClient.SendEmergencyAlert(ctx, em.PlantID, em.Title, em.Severity)
	}

	_ = s.publisher.Publish(ctx, events.EventEmergencyDeclared, em)
	prahariLogger.Warn(ctx, "Industrial emergency declared!", prahariLogger.String("emergency_number", em.EmergencyNumber))
	return nil
}

func (s *Service) ActivateResponse(ctx context.Context, id string) error {
	em, err := s.repo.GetEmergencyByID(ctx, id)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(em.Status), status.CodeResponseActivated); err != nil {
		return err
	}

	em.Status = string(status.CodeResponseActivated)
	now := time.Now()
	em.CommandEstablishedAt = &now
	em.UpdatedAt = now

	if err := s.repo.SaveEmergency(ctx, em); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventEmergencyResponseActivate, em)
	prahariLogger.Info(ctx, "Emergency response activated and ICS command established", prahariLogger.String("id", id))
	return nil
}

func (s *Service) DeployResource(ctx context.Context, emergencyID, resourceID string, qty int) error {
	if err := s.repo.DeployResource(ctx, resourceID, qty); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventResourceDeployed, map[string]any{
		"emergency_id": emergencyID,
		"resource_id":  resourceID,
		"quantity":     qty,
	})
	prahariLogger.Info(ctx, "Emergency resource deployed", prahariLogger.String("resource_id", resourceID))
	return nil
}

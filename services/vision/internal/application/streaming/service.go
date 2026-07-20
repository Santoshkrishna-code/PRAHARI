package streaming

import (
	"context"
	"fmt"
	"time"

	"prahari/services/vision/internal/domain/camera"
	"prahari/services/vision/internal/domain/events"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveCamera(ctx context.Context, c *camera.Camera) error
	GetCameraByID(ctx context.Context, id string) (*camera.Camera, error)
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type Service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, pub EventPublisher) *Service {
	return &Service{repo: repo, publisher: pub}
}

func (s *Service) RegisterCamera(ctx context.Context, c *camera.Camera) error {
	c.ID = fmt.Sprintf("cam-%d", time.Now().UnixNano())
	c.Status = "ONLINE"
	c.CreatedAt = time.Now()

	if err := s.repo.SaveCamera(ctx, c); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "New camera IP registered successfully",
		prahariLogger.String("camera_id", c.ID),
		prahariLogger.String("ip", c.IPAddress))
	return nil
}

func (s *Service) DisconnectCamera(ctx context.Context, id string) error {
	c, err := s.repo.GetCameraByID(ctx, id)
	if err != nil {
		return err
	}

	c.Status = "OFFLINE"
	if err := s.repo.SaveCamera(ctx, c); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventVisionCameraOffline, c)
	return nil
}

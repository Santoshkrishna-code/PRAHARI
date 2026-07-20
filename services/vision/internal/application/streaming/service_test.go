package streaming_test

import (
	"context"
	"testing"

	"prahari/services/vision/internal/application/streaming"
	"prahari/services/vision/internal/domain/camera"
)

type mockRepo struct {
	savedCamera *camera.Camera
}

func (m *mockRepo) SaveCamera(ctx context.Context, c *camera.Camera) error {
	m.savedCamera = c
	return nil
}

func (m *mockRepo) GetCameraByID(ctx context.Context, id string) (*camera.Camera, error) {
	return m.savedCamera, nil
}

type mockPublisher struct{}

func (m *mockPublisher) Publish(ctx context.Context, eventType string, payload any) error {
	return nil
}

func TestRegisterCamera(t *testing.T) {
	repo := &mockRepo{}
	svc := streaming.NewService(repo, &mockPublisher{})

	cam := &camera.Camera{
		PlantID:   "P01",
		Name:      "Zone A Gate",
		IPAddress: "10.0.0.12",
	}

	err := svc.RegisterCamera(context.Background(), cam)
	if err != nil {
		t.Fatalf("unexpected camera registration error: %v", err)
	}

	if cam.ID == "" {
		t.Error("expected generated camera ID")
	}

	if cam.Status != "ONLINE" {
		t.Errorf("expected state ONLINE, got %s", cam.Status)
	}
}

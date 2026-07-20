package synchronization_test

import (
	"context"
	"testing"

	"prahari/services/digitaltwin/internal/application/synchronization"
	"prahari/services/digitaltwin/internal/domain/state"
)

type mockRepo struct {
	savedState *state.LiveState
}

func (m *mockRepo) SaveLiveState(ctx context.Context, s *state.LiveState) error {
	m.savedState = s
	return nil
}

type mockPublisher struct {
	publishedType string
}

func (m *mockPublisher) Publish(ctx context.Context, eventType string, payload any) error {
	m.publishedType = eventType
	return nil
}

func TestSyncTelemetry(t *testing.T) {
	repo := &mockRepo{}
	pub := &mockPublisher{}
	svc := synchronization.NewService(repo, pub)

	// Sync valid telemetry quality
	err := svc.SyncTelemetry(context.Background(), "twin-1", "eq-1", 45.2, "GOOD")
	if err != nil {
		t.Fatalf("unexpected sync error: %v", err)
	}

	if repo.savedState == nil {
		t.Fatal("expected state payload to be saved to database")
	}

	if repo.savedState.Value != 45.2 {
		t.Errorf("expected value 45.2, got %f", repo.savedState.Value)
	}

	if pub.publishedType != "digitaltwin.state.updated" {
		t.Errorf("expected event type digitaltwin.state.updated, got %s", pub.publishedType)
	}

	// Try syncing bad quality
	repo.savedState = nil
	err = svc.SyncTelemetry(context.Background(), "twin-1", "eq-1", 99.9, "BAD")
	if err != nil {
		t.Fatalf("unexpected error on bad quality sync: %v", err)
	}

	if repo.savedState != nil {
		t.Error("expected state updates to be skipped when quality is BAD")
	}
}

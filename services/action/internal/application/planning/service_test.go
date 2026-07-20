package planning_test

import (
	"context"
	"testing"
	"time"

	"prahari/services/action/internal/application/planning"
	"prahari/services/action/internal/domain/action"
)

type mockRepo struct{}

func (m *mockRepo) SaveAction(ctx context.Context, act *action.Action) error {
	return nil
}

type mockPublisher struct{}

func (m *mockPublisher) Publish(ctx context.Context, eventType string, payload any) error {
	return nil
}

func TestCreateActionItem(t *testing.T) {
	svc := planning.NewService(&mockRepo{}, &mockPublisher{})

	act := &action.Action{
		PlantID:      "PLANT-01",
		SourceModule: "INCIDENT",
		SourceRefID:  "INC-001",
		Title:        "Replace corroded safety valve",
		ActionType:   "CORRECTIVE",
		DueDate:      time.Now().Add(7 * 24 * time.Hour),
	}

	err := svc.CreateActionItem(context.Background(), act)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if act.ID == "" {
		t.Error("expected action ID to be populated")
	}
	if act.Status != "CREATED" {
		t.Errorf("expected status CREATED, got %s", act.Status)
	}
}

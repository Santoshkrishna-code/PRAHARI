package scheduling_test

import (
	"context"
	"testing"
	"time"

	"prahari/services/meetings/internal/application/scheduling"
	"prahari/services/meetings/internal/domain/meeting"
)

type mockRepo struct{}

func (m *mockRepo) SaveMeeting(ctx context.Context, mtg *meeting.Meeting) error {
	return nil
}

type mockPublisher struct{}

func (m *mockPublisher) Publish(ctx context.Context, eventType string, payload any) error {
	return nil
}

func TestScheduleMeeting(t *testing.T) {
	svc := scheduling.NewService(&mockRepo{}, &mockPublisher{})

	mtg := &meeting.Meeting{
		PlantID:     "PLANT-01",
		MeetingType: "TOOLBOX_TALK",
		Title:       "Morning Shift Safety Briefing",
		Location:    "Assembly Area B",
		ScheduledAt: time.Now().Add(2 * time.Hour),
		OrganizerID: "safety-officer-001",
	}

	err := svc.ScheduleMeeting(context.Background(), mtg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if mtg.ID == "" {
		t.Error("expected meeting ID to be populated")
	}
	if mtg.Status != "PLANNED" {
		t.Errorf("expected status PLANNED, got %s", mtg.Status)
	}
}

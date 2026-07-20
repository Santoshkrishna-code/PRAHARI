package events_test

import (
	"testing"

	"prahari/shared/kafka/events"
)

type WorkerAlert struct {
	WorkerID string `json:"worker_id"`
	Reason   string `json:"reason"`
}

func TestEvent_WrapUnwrap(t *testing.T) {
	payload := WorkerAlert{WorkerID: "w-10", Reason: "heart_rate_spike"}

	envelopeBytes, err := events.WrapEvent("id-123", "WorkerHeartRateSpike", "worker-service", "1.0.1", payload)
	if err != nil {
		t.Fatalf("failed to wrap event: %v", err)
	}

	var parsed WorkerAlert
	env, err := events.UnwrapEvent(envelopeBytes, &parsed)
	if err != nil {
		t.Fatalf("failed to unwrap event: %v", err)
	}

	if env.ID != "id-123" || env.Type != "WorkerHeartRateSpike" || env.Version != "1.0.1" {
		t.Errorf("envelope headers mismatch, got: %+v", env)
	}

	if parsed.WorkerID != payload.WorkerID || parsed.Reason != payload.Reason {
		t.Errorf("envelope payload data mismatch, got: %+v", parsed)
	}
}

func TestVersioning_IsCompatible(t *testing.T) {
	if !events.IsCompatible("1.0.0", "1.2.3") {
		t.Error("expected 1.0.0 and 1.2.3 to be compatible")
	}

	if events.IsCompatible("1.0.0", "2.0.0") {
		t.Error("expected 1.0.0 and 2.0.0 to be incompatible due to major change")
	}
}

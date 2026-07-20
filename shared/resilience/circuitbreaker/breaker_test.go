package circuitbreaker_test

import (
	"errors"
	"testing"
	"time"

	"prahari/shared/resilience/circuitbreaker"
)

func TestCircuitBreaker_StateTransitions(t *testing.T) {
	// Set failure ratio to 0.5 (50%), cooldown to 50ms, trips after 5 requests
	cb := circuitbreaker.NewCircuitBreaker("test-cb", 0.5, 50*time.Millisecond)

	if cb.GetState() != circuitbreaker.StateClosed {
		t.Errorf("expected initially CLOSED, got %s", cb.GetState())
	}

	// Trigger 5 errors to trip the circuit
	for i := 0; i < 5; i++ {
		_, _ = cb.Execute(func() (interface{}, error) {
			return nil, errors.New("network failure")
		})
	}

	// State should now transition to OPEN
	if cb.GetState() != circuitbreaker.StateOpen {
		t.Errorf("expected state to transition to OPEN, got %s", cb.GetState())
	}

	// Subsequent runs should block immediately with ErrCircuitOpen
	_, err := cb.Execute(func() (interface{}, error) {
		return "success", nil
	})
	if !errors.Is(err, circuitbreaker.ErrCircuitOpen) {
		t.Errorf("expected ErrCircuitOpen, got %v", err)
	}

	// Wait for cooldown period (50ms) to allow transition to HALF-OPEN
	time.Sleep(60 * time.Millisecond)

	if cb.GetState() != circuitbreaker.StateHalfOpen {
		t.Errorf("expected state to transition to HALF-OPEN after cooldown, got %s", cb.GetState())
	}

	// Run successful call to close the circuit again
	_, err = cb.Execute(func() (interface{}, error) {
		return "success", nil
	})
	if err != nil {
		t.Fatalf("expected successful call to close circuit, got: %v", err)
	}

	if cb.GetState() != circuitbreaker.StateClosed {
		t.Errorf("expected state to transition back to CLOSED on success, got %s", cb.GetState())
	}
}

package fallback_test

import (
	"errors"
	"testing"

	"prahari/shared/resilience/fallback"
)

func TestFallback_Execute(t *testing.T) {
	primaryErr := errors.New("primary database connection failed")

	// Case 1: Primary succeeds -> expect fallback bypassed
	res, err := fallback.Execute(
		func() (interface{}, error) { return "primary-data", nil },
		func(err error) (interface{}, error) { return "fallback-data", nil },
	)

	if err != nil {
		t.Fatalf("expected successful primary run, got: %v", err)
	}

	if res.(string) != "primary-data" {
		t.Errorf("expected 'primary-data', got '%s'", res.(string))
	}

	// Case 2: Primary fails -> expect fallback execution run
	res, err = fallback.Execute(
		func() (interface{}, error) { return nil, primaryErr },
		func(err error) (interface{}, error) {
			if errors.Is(err, primaryErr) {
				return "fallback-data", nil
			}
			return nil, err
		},
	)

	if err != nil {
		t.Fatalf("expected fallback execution to succeed, got: %v", err)
	}

	if res.(string) != "fallback-data" {
		t.Errorf("expected 'fallback-data', got '%s'", res.(string))
	}
}

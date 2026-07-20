package health_test

import (
	"context"
	"errors"
	"testing"

	prahariRes "prahari/shared/resilience"
	"prahari/shared/resilience/health"
)

type MockChecker struct {
	name string
	err  error
}

func (m *MockChecker) Name() string { return m.name }
func (m *MockChecker) Ping(ctx context.Context) error { return m.err }

func TestHealthRegistry(t *testing.T) {
	c1 := &MockChecker{name: "PostgreSQL", err: nil}
	c2 := &MockChecker{name: "Redis", err: errors.New("connection timeout")}

	reg := health.NewRegistry(c1, c2)
	results := reg.CheckHealth(context.Background())

	if err, ok := results["PostgreSQL"]; !ok || err != nil {
		t.Errorf("expected PostgreSQL to be ONLINE (nil error), got: %v", err)
	}

	if err, ok := results["Redis"]; !ok || err == nil {
		t.Errorf("expected Redis to be OFFLINE (timeout error), got nil")
	}
}

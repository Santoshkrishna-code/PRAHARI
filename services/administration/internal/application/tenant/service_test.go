package tenant_test

import (
	"context"
	"testing"

	"prahari/services/administration/internal/application/tenant"
	tenantDomain "prahari/services/administration/internal/domain/tenant"
)

type mockRepo struct {
	savedTenant *tenantDomain.Tenant
}

func (m *mockRepo) SaveTenant(ctx context.Context, t *tenantDomain.Tenant) error {
	m.savedTenant = t
	return nil
}

type mockPublisher struct{}

func (m *mockPublisher) Publish(ctx context.Context, eventType string, payload any) error {
	return nil
}

func TestCreateTenant(t *testing.T) {
	repo := &mockRepo{}
	svc := tenant.NewService(repo, &mockPublisher{})

	newTenant := &tenantDomain.Tenant{
		Name:   "Astra Space",
		Domain: "astra.io",
	}

	err := svc.CreateTenant(context.Background(), newTenant)
	if err != nil {
		t.Fatalf("unexpected error during tenant creation: %v", err)
	}

	if newTenant.ID == "" {
		t.Error("expected generated tenant ID to be non-empty")
	}

	if newTenant.Status != "ACTIVE" {
		t.Errorf("expected tenant status to be ACTIVE, got %s", newTenant.Status)
	}
}

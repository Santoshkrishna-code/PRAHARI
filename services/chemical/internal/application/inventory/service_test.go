package inventory_test

import (
	"context"
	"testing"

	"prahari/services/chemical/internal/application/inventory"
	"prahari/services/chemical/internal/domain/container"
)

type mockRepo struct {
	con *container.Container
}

func (m *mockRepo) SaveContainer(ctx context.Context, con *container.Container) error {
	m.con = con
	return nil
}

func (m *mockRepo) GetContainerByID(ctx context.Context, id string) (*container.Container, error) {
	return m.con, nil
}

type mockPublisher struct{}

func (m *mockPublisher) Publish(ctx context.Context, eventType string, payload any) error {
	return nil
}

func TestReceiveContainer(t *testing.T) {
	repo := &mockRepo{}
	svc := inventory.NewService(repo, &mockPublisher{})

	con := &container.Container{
		ChemicalID: "chem-01",
		Barcode:    "BC-777",
		Capacity:   10.0,
	}

	err := svc.ReceiveContainer(context.Background(), con)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if con.ID == "" {
		t.Error("expected generated container ID")
	}
	if con.Status != "RECEIVED" {
		t.Errorf("expected status RECEIVED, got %s", con.Status)
	}
}

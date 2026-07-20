package connectors_test

import (
	"context"
	"testing"

	"prahari/services/integration/internal/application/connectors"
	"prahari/services/integration/internal/domain/connector"
)

type mockRepo struct {
	conn *connector.Connector
}

func (m *mockRepo) SaveConnector(ctx context.Context, c *connector.Connector) error {
	m.conn = c
	return nil
}

func (m *mockRepo) GetConnectorByID(ctx context.Context, id string) (*connector.Connector, error) {
	return m.conn, nil
}

type mockPublisher struct{}

func (m *mockPublisher) Publish(ctx context.Context, eventType string, payload any) error {
	return nil
}

func TestRegisterConnector(t *testing.T) {
	repo := &mockRepo{}
	svc := connectors.NewService(repo, &mockPublisher{})

	conn := &connector.Connector{
		Name: "Maximo Prod API",
		Type: "MAXIMO",
	}

	err := svc.RegisterConnector(context.Background(), conn)
	if err != nil {
		t.Fatalf("unexpected registration error: %v", err)
	}

	if conn.ID == "" {
		t.Error("expected generated connector ID")
	}

	if conn.Status != "DISCONNECTED" {
		t.Errorf("expected state DISCONNECTED, got %s", conn.Status)
	}
}

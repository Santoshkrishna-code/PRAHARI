package fixtures_test

import (
	"testing"

	"prahari/shared/testing/fixtures"
)

func TestFixturesCreation(t *testing.T) {
	admin := fixtures.NewAdminClaims()
	if admin.UserID != "usr-admin-99" || admin.Role != "Admin" {
		t.Errorf("mismatched admin fixture parameters: %+v", admin)
	}

	worker := fixtures.NewWorkerClaims()
	if worker.UserID != "usr-worker-11" || worker.Role != "Worker" {
		t.Errorf("mismatched worker fixture parameters: %+v", worker)
	}

	payload := fixtures.NewIncidentPayload()
	if len(payload) == 0 {
		t.Error("expected positive JSON template array, got empty")
	}
}

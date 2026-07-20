package policy

import (
	"testing"
	"time"

	"prahari/services/moc/internal/domain/changerequest"
)

func TestRequiresFullHazardAnalysis(t *testing.T) {
	reqHighRisk := &changerequest.Request{RiskLevel: "HIGH", ChangeType: "MECHANICAL"}
	if !RequiresFullHazardAnalysis(reqHighRisk) {
		t.Errorf("Expected HIGH risk request to require full hazard analysis")
	}

	reqProcess := &changerequest.Request{RiskLevel: "LOW", ChangeType: "PROCESS"}
	if !RequiresFullHazardAnalysis(reqProcess) {
		t.Errorf("Expected PROCESS change type to require full hazard analysis")
	}

	reqLowMechanical := &changerequest.Request{RiskLevel: "LOW", ChangeType: "DOCUMENTATION"}
	if RequiresFullHazardAnalysis(reqLowMechanical) {
		t.Errorf("Expected LOW documentation change to NOT require full hazard analysis")
	}
}

func TestIsTemporaryExpired(t *testing.T) {
	past := time.Now().Add(-24 * time.Hour)
	reqExpired := &changerequest.Request{
		Category:   changerequest.CategoryTemporary,
		ExpiryDate: &past,
		UpdatedAt:  time.Now(),
	}
	if !IsTemporaryExpired(reqExpired) {
		t.Errorf("Expected temporary change with past expiry date to be expired")
	}
}

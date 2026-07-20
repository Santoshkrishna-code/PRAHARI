package policy

import (
	"testing"
)

func TestCalculateRiskRank(t *testing.T) {
	rank, cat := CalculateRiskRank(5, 4)
	if rank != 20 || cat != "UNACCEPTABLE" {
		t.Errorf("CalculateRiskRank(5,4) got rank=%d, cat=%s; want 20, UNACCEPTABLE", rank, cat)
	}

	rankLow, catLow := CalculateRiskRank(2, 2)
	if rankLow != 4 || catLow != "LOW" {
		t.Errorf("CalculateRiskRank(2,2) got rank=%d, cat=%s; want 4, LOW", rankLow, catLow)
	}
}

func TestCalculateLOPARequiredSIL(t *testing.T) {
	// Initiating freq 1e-1 (0.1/yr), Target 1e-5 (/yr), IPL mitigation 1e-2 (0.01)
	mitFreq, rrf, sil := CalculateLOPARequiredSIL(0.1, 1e-5, 0.01)
	if sil != "SIL-2" {
		t.Errorf("CalculateLOPARequiredSIL() got sil=%s, rrf=%.2f, mitFreq=%e; want SIL-2", sil, rrf, mitFreq)
	}
}

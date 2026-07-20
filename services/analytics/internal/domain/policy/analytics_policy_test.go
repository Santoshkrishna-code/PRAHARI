package policy_test

import (
	"testing"

	"prahari/services/analytics/internal/domain/alert"
	"prahari/services/analytics/internal/domain/kpi"
	"prahari/services/analytics/internal/domain/policy"
)

func TestEvaluateKPIStatus(t *testing.T) {
	s1 := policy.EvaluateKPIStatus(100.0, 105.0)
	if s1 != "ON_TRACK" {
		t.Errorf("expected ON_TRACK, got %s", s1)
	}

	s2 := policy.EvaluateKPIStatus(100.0, 85.0)
	if s2 != "AT_RISK" {
		t.Errorf("expected AT_RISK, got %s", s2)
	}

	s3 := policy.EvaluateKPIStatus(100.0, 70.0)
	if s3 != "CRITICAL" {
		t.Errorf("expected CRITICAL, got %s", s3)
	}
}

func TestIsAlertTriggered(t *testing.T) {
	rule := &alert.Rule{
		Threshold: 50.0,
		Operator:  "GREATER_THAN",
		Active:    true,
	}

	if !policy.IsAlertTriggered(rule, 55.0) {
		t.Error("expected alert to be triggered when value exceeds threshold")
	}

	if policy.IsAlertTriggered(rule, 45.0) {
		t.Error("expected alert not to be triggered when value is below threshold")
	}
}

func TestEvaluateKPIThreshold(t *testing.T) {
	kpiOk := &kpi.KPI{TargetVal: 10.0, ActualVal: 8.0}
	kpiViolated := &kpi.KPI{TargetVal: 10.0, ActualVal: 4.0}

	if policy.EvaluateKPIThreshold(kpiOk) {
		t.Error("expected KPI threshold not violated")
	}

	if !policy.EvaluateKPIThreshold(kpiViolated) {
		t.Error("expected KPI threshold violated")
	}
}

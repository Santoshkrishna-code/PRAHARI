package policy

import (
	"prahari/services/analytics/internal/domain/alert"
	"prahari/services/analytics/internal/domain/kpi"
)

// EvaluateKPIStatus determines KPI status based on current vs target values.
func EvaluateKPIStatus(target, actual float64) string {
	if actual >= target {
		return "ON_TRACK"
	}
	if actual >= target*0.8 {
		return "AT_RISK"
	}
	return "CRITICAL"
}

// IsAlertTriggered checks if a metric value exceeds defined alert threshold operator criteria.
func IsAlertTriggered(rule *alert.Rule, value float64) bool {
	if rule == nil || !rule.Active {
		return false
	}
	switch rule.Operator {
	case "GREATER_THAN":
		return value > rule.Threshold
	case "LESS_THAN":
		return value < rule.Threshold
	case "EQUALS":
		return value == rule.Threshold
	default:
		return false
	}
}

// EvaluateKPIThreshold returns true if the KPI actual level has violated safety thresholds.
func EvaluateKPIThreshold(k *kpi.KPI) bool {
	if k == nil {
		return false
	}
	return k.ActualVal < k.TargetVal*0.5
}

package policy

import (
	"prahari/services/moc/internal/domain/changerequest"
)

// RequiresFullHazardAnalysis checks if a change request requires full PHA/HAZOP due to risk level or type.
func RequiresFullHazardAnalysis(req *changerequest.Request) bool {
	if req.RiskLevel == "HIGH" || req.RiskLevel == "CRITICAL" {
		return true
	}
	if req.ChangeType == "PROCESS" || req.ChangeType == "CHEMICAL" || req.ChangeType == "AUTOMATION" {
		return true
	}
	return false
}

// IsTemporaryExpired checks if a temporary change has surpassed its expiration date.
func IsTemporaryExpired(req *changerequest.Request) bool {
	if req.Category != changerequest.CategoryTemporary {
		return false
	}
	if req.ExpiryDate == nil {
		return false
	}
	return req.ExpiryDate.Before(req.UpdatedAt)
}

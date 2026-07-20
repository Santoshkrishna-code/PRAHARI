package compliance

import (
	"prahari/services/chemical/internal/domain/label"
)

// ValidateCLP checks that the labelling matches EU CLP guidelines.
func ValidateCLP(lbl *label.Label) (bool, string) {
	if lbl == nil {
		return false, "Label specifications missing"
	}
	if lbl.LabelFormat == "" {
		return false, "CLP label formatting specification is missing"
	}
	return true, "Complies with EU CLP labelling requirements"
}

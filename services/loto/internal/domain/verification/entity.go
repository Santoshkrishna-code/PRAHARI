package verification

import "time"

// ZeroEnergy tracks the physical verification testing (e.g. pressure check, voltage test) confirming zero energy.
type ZeroEnergy struct {
	ID             string    `json:"id"`
	CertificateID  string    `json:"certificate_id"`
	VerifiedBy     string    `json:"verified_by"`
	VerificationAt time.Time `json:"verification_at"`
	TestPassed     bool      `json:"test_passed"`
	TestMethod     string    `json:"test_method"` // TRY_START, PRESSURE_GAUGE, VOLTAGE_METER
	Notes          string    `json:"notes"`
}

package electronicsignature

import "time"

// Signature represents 21 CFR Part 11 / eIDAS compliant electronic signature sign-off.
type Signature struct {
	ID            string    `json:"id"`
	DocumentID    string    `json:"document_id"`
	VersionID     string    `json:"version_id"`
	SignerID      string    `json:"signer_id"`
	SignerRole    string    `json:"signer_role"`
	SignatureHash string    `json:"signature_hash"`
	SignedAt      time.Time `json:"signed_at"`
}

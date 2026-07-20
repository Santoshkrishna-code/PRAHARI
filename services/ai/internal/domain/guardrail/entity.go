package guardrail

import "time"

// PolicyRule defines PII redact flags and input validation blocks.
type PolicyRule struct {
	ID         string    `json:"id"`
	RuleName   string    `json:"rule_name"` // E.g., redact_pii, block_jailbreak
	ActionType string    `json:"action_type"` // REDACT, BLOCK
	Active     bool      `json:"active"`
	UpdatedAt  time.Time `json:"updated_at"`
}

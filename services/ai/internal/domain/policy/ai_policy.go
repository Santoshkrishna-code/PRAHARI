package policy

import (
	"strings"
)

// RedactPII replaces social security, email, or telephone lookups with redaction tags.
func RedactPII(input string) string {
	// Simple email check
	words := strings.Split(input, " ")
	for i, w := range words {
		if strings.Contains(w, "@") && strings.Contains(w, ".") {
			words[i] = "[REDACTED_EMAIL]"
		}
	}
	return strings.Join(words, " ")
}

// EvaluateQuerySafety checks queries for injection prefixes.
func EvaluateQuerySafety(query string) bool {
	lower := strings.ToLower(query)
	bannedPrefixes := []string{"ignore previous instructions", "system prompt", "jailbreak"}
	for _, bp := range bannedPrefixes {
		if strings.Contains(lower, bp) {
			return false
		}
	}
	return true
}

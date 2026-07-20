package hashing

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	hasUpper   = regexp.MustCompile(`[A-Z]`)
	hasLower   = regexp.MustCompile(`[a-z]`)
	hasNumber  = regexp.MustCompile(`[0-9]`)
	hasSpecial = regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};':",\./<>\?~` + "`" + `]`)
)

// VerifyPasswordStrength checks that the password meets minimum complexity constraints.
func VerifyPasswordStrength(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if !hasUpper.MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower.MatchString(password) {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasNumber.MatchString(password) {
		return errors.New("password must contain at least one digit")
	}
	if !hasSpecial.MatchString(password) {
		return errors.New("password must contain at least one special character")
	}
	return nil
}

// DetectBcryptUpgrade checks if an existing bcrypt hash's work factor is lower than the targetCost.
func DetectBcryptUpgrade(hash string, targetCost int) bool {
	// Bcrypt prefix matching $2a$, $2b$, or $2y$
	if !strings.HasPrefix(hash, "$2") {
		return false
	}
	
	parts := strings.Split(hash, "$")
	if len(parts) < 3 {
		return false
	}
	
	cost, err := strconv.Atoi(parts[2])
	if err != nil {
		return false
	}
	
	return cost < targetCost
}

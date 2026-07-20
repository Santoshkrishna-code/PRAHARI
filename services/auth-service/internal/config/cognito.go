package config

import (
	"fmt"
)

type CognitoConfig struct {
	UserPoolID string `env:"COGNITO_USER_POOL_ID"`
	ClientID   string `env:"COGNITO_CLIENT_ID"`
}

// GetJWKSURL returns the standard Cognito JSON Web Key Set url for token verification
func (c *CognitoConfig) GetJWKSURL(region string) string {
	return fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, c.UserPoolID)
}

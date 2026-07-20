package domain

// TokenPair encapsulates the access, id, and refresh JWTs returned by the identity provider.
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int32  `json:"expires_in"`
}

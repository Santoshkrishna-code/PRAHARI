package permission

// Permission defines atomic permission actions scopes.
type Permission struct {
	Scope       string `json:"scope"`
	Description string `json:"description"`
}

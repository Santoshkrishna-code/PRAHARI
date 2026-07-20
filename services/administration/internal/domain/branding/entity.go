package branding

// Config defines tenant specific appearance customisations.
type Config struct {
	ID             string `json:"id"`
	TenantID       string `json:"tenant_id"`
	LogoURL        string `json:"logo_url,omitempty"`
	PrimaryColor   string `json:"primary_color"`
	SecondaryColor string `json:"secondary_color"`
	CustomCSS      string `json:"custom_css,omitempty"`
}

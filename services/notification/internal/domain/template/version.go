package template

// Version maps version configurations tracking template evolution.
type Version struct {
	TemplateID string `json:"template_id"`
	VersionNo  int    `json:"version_no"`
	Active     bool   `json:"active"`
}

package aws

// Config holds target AWS credentials profiles, regions, and LocalStack endpoints options.
type Config struct {
	Region      string `json:"region"`
	Profile     string `json:"profile"`
	RoleARN     string `json:"role_arn"`
	EndpointURL string `json:"endpoint_url"` // Custom endpoint url resolving to LocalStack
}

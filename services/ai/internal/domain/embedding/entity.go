package embedding

// Config represents vectors generation configurations.
type Config struct {
	ModelName string `json:"model_name"`
	Dimension int    `json:"dimension"`
}

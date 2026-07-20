package transformation

// Pipeline defines custom scripting or transform format instructions.
type Pipeline struct {
	ID             string `json:"id"`
	ConnectorID    string `json:"connector_id"`
	InputFormat    string `json:"input_format"`  // JSON, XML, CSV
	OutputFormat   string `json:"output_format"` // JSON, XML, CSV
	TransformScript string `json:"transform_script"`
}

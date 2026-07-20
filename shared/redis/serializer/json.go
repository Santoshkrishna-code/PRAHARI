package serializer

import (
	"encoding/json"
)

// JSONSerializer implements redis.Serializer interface using JSON format.
type JSONSerializer struct{}

// NewJSONSerializer constructs a JSON serializer.
func NewJSONSerializer() *JSONSerializer {
	return &JSONSerializer{}
}

// Marshal encodes Go objects into JSON byte slices.
func (s *JSONSerializer) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal decodes JSON bytes slices into target Go pointers.
func (s *JSONSerializer) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

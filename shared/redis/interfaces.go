package redis

// Serializer defines the interface for encoding/decoding cache entries.
type Serializer interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

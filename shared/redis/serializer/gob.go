package serializer

import (
	"bytes"
	"encoding/gob"
)

// GOBSerializer implements redis.Serializer interface using Go binary GOB format.
type GOBSerializer struct{}

// NewGOBSerializer constructs a GOB serializer.
func NewGOBSerializer() *GOBSerializer {
	return &GOBSerializer{}
}

// Marshal encodes Go objects into binary GOB byte slices.
func (s *GOBSerializer) Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Unmarshal decodes binary GOB bytes slices into target Go pointers.
func (s *GOBSerializer) Unmarshal(data []byte, v interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(v)
}

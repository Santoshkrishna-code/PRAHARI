package events

import (
	"encoding/json"
	"fmt"
	"time"
)

// WrapEvent marshals a payload and packages it into an Envelope, returning the serialized envelope bytes.
func WrapEvent(id, eventType, source, version string, payload interface{}) ([]byte, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal event payload: %w", err)
	}

	env := Envelope{
		ID:        id,
		Type:      eventType,
		Source:    source,
		Timestamp: time.Now(),
		Version:   version,
		Payload:   payloadBytes,
	}

	envelopeBytes, err := json.Marshal(env)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal event envelope: %w", err)
	}

	return envelopeBytes, nil
}

// UnwrapEvent unmarshals data into an Envelope and returns the envelope, allowing callers to decode the inner payload.
func UnwrapEvent(data []byte, target interface{}) (*Envelope, error) {
	var env Envelope
	err := json.Unmarshal(data, &env)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal event envelope: %w", err)
	}

	err = json.Unmarshal(env.Payload, target)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal inner event payload: %w", err)
	}

	return &env, nil
}

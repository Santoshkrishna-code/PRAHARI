package schema

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// SchemaValidator enforces structured schema checks on incoming/outgoing event payloads.
type SchemaValidator struct {
	val *validator.Validate
}

// NewSchemaValidator constructs a SchemaValidator.
func NewSchemaValidator() *SchemaValidator {
	return &SchemaValidator{val: validator.New()}
}

// Validate checks that the event payload struct conforms to play-ground validation tags.
func (v *SchemaValidator) Validate(payload interface{}) error {
	err := v.val.Struct(payload)
	if err != nil {
		return fmt.Errorf("event schema constraints check failed: %w", err)
	}
	return nil
}

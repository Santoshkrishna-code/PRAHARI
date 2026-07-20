package schema_test

import (
	"testing"

	"prahari/shared/kafka/schema"
)

type TestPayload struct {
	Email string `validate:"required,email"`
	Port  int    `validate:"required,gte=80"`
}

func TestSchemaValidator(t *testing.T) {
	v := schema.NewSchemaValidator()

	// Valid payload
	p1 := TestPayload{Email: "worker@prahari.gov", Port: 8080}
	err := v.Validate(p1)
	if err != nil {
		t.Fatalf("expected valid payload to pass, got: %v", err)
	}

	// Invalid payload (invalid email, port out of range)
	p2 := TestPayload{Email: "invalid-email", Port: 79}
	err = v.Validate(p2)
	if err == nil {
		t.Fatal("expected validation check to error on invalid values, got nil")
	}
}

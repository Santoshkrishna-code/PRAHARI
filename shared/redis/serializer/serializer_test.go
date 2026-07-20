package serializer_test

import (
	"encoding/gob"
	"testing"

	"prahari/shared/redis/serializer"
)

type User struct {
	Name string
	Age  int
}

func init() {
	// Register type for GOB parser serialization
	gob.Register(User{})
}

func TestJSONSerializer(t *testing.T) {
	s := serializer.NewJSONSerializer()
	user := User{Name: "Santosh", Age: 30}

	data, err := s.Marshal(user)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var parsed User
	err = s.Unmarshal(data, &parsed)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if parsed.Name != user.Name || parsed.Age != user.Age {
		t.Errorf("expected %+v, got %+v", user, parsed)
	}
}

func TestGOBSerializer(t *testing.T) {
	s := serializer.NewGOBSerializer()
	user := User{Name: "Santosh", Age: 30}

	data, err := s.Marshal(user)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var parsed User
	err = s.Unmarshal(data, &parsed)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if parsed.Name != user.Name || parsed.Age != user.Age {
		t.Errorf("expected %+v, got %+v", user, parsed)
	}
}

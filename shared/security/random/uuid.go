package random

import (
	"fmt"
)

// GenerateUUIDv4 returns a cryptographically secure UUID version 4 string conforming to RFC 4122.
func GenerateUUIDv4() (string, error) {
	uuid, err := Bytes(16)
	if err != nil {
		return "", err
	}

	// Set version 4 bits: 0100xxxx
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	
	// Set variant bits: 10xxxxxx
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	return fmt.Sprintf(
		"%x-%x-%x-%x-%x",
		uuid[0:4],
		uuid[4:6],
		uuid[6:8],
		uuid[8:10],
		uuid[10:],
	), nil
}

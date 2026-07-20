package security

// Config holds runtime options for encryption, hashing, and token signatures.
type Config struct {
	AESKey           []byte `json:"-"` // Redacted in logs, holds raw 32-byte key
	BcryptCost       int    `json:"bcrypt_cost"`
	Argon2Memory     uint32 `json:"argon2_memory_kb"`
	Argon2Iterations uint32 `json:"argon2_iterations"`
	Argon2Parallelism uint8  `json:"argon2_parallelism"`
	JWTIssuer        string `json:"jwt_issuer"`
	JWTAudience      string `json:"jwt_audience"`
}

// DefaultConfig returns safe production defaults.
func DefaultConfig() Config {
	return Config{
		BcryptCost:        12,
		Argon2Memory:      65536, // 64 MB memory
		Argon2Iterations:  3,
		Argon2Parallelism: 4,
	}
}

package security

import "context"

// Encrypter defines structural interface for cryptographic encrypt operations.
type Encrypter interface {
	Encrypt(ctx context.Context, plaintext []byte) ([]byte, error)
}

// Decrypter defines structural interface for cryptographic decrypt operations.
type Decrypter interface {
	Decrypt(ctx context.Context, ciphertext []byte) ([]byte, error)
}

// Cipher compiles both encryption and decryption capabilities.
type Cipher interface {
	Encrypter
	Decrypter
}

// Hasher defines structural interface for hashing passwords and comparing hashes.
type Hasher interface {
	Hash(ctx context.Context, data string) (string, error)
	Verify(ctx context.Context, data, hash string) (bool, error)
}

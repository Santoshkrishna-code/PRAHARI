package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// GenerateRSAKeys generates a new RSA private and public key pair.
func GenerateRSAKeys(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate rsa key: %w", err)
	}
	return privateKey, &privateKey.PublicKey, nil
}

// ExportPrivateKeyPEM serializes an RSA private key to PEM block bytes.
func ExportPrivateKeyPEM(priv *rsa.PrivateKey) []byte {
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priv),
	}
	return pem.EncodeToMemory(block)
}

// ExportPublicKeyPEM serializes an RSA public key to PEM block bytes.
func ExportPublicKeyPEM(pub *rsa.PublicKey) ([]byte, error) {
	derBytes, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, err
	}
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derBytes,
	}
	return pem.EncodeToMemory(block), nil
}

// ParsePrivateKeyPEM parses PKCS1 PEM private keys.
func ParsePrivateKeyPEM(pemBytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing RSA Private Key")
	}
	
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse rsa private key: %w", err)
	}
	return priv, nil
}

// ParsePublicKeyPEM parses PKIX PEM public keys.
func ParsePublicKeyPEM(pemBytes []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing RSA Public Key")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse rsa public key: %w", err)
	}

	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("PEM block did not contain valid RSA public key type")
	}

	return pub, nil
}

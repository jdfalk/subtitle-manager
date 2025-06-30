// file: pkg/backups/encryption.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174003

package backups

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// AESEncryption implements the Encryption interface using AES-GCM.
type AESEncryption struct {
	key []byte
}

// NewAESEncryption creates a new AES encryption instance.
// The key should be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256.
func NewAESEncryption(key []byte) (*AESEncryption, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, fmt.Errorf("invalid key size: %d (must be 16, 24, or 32 bytes)", len(key))
	}
	return &AESEncryption{key: key}, nil
}

// Encrypt encrypts the input data using AES-GCM.
func (ae *AESEncryption) Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(ae.key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

// Decrypt decrypts the input data using AES-GCM.
func (ae *AESEncryption) Decrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(ae.key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt data: %w", err)
	}

	return plaintext, nil
}
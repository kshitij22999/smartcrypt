package algo

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// AESGCMEncryptor implements the Encryptor interface using AES-GCM.
type AESGCMEncryptor struct {
	key []byte // 16, 24, or 32 bytes for AES-128/192/256
}

// NewAESGCMEncryptor creates a new AESGCMEncryptor with a provided key.
func NewAESGCMEncryptor(key []byte) (*AESGCMEncryptor, error) {
	keyLen := len(key)
	if keyLen != 16 && keyLen != 24 && keyLen != 32 {
		return nil, errors.New("invalid AES key size (must be 16, 24, or 32 bytes)")
	}
	return &AESGCMEncryptor{key: key}, nil
}

// Encrypt encrypts plaintext using AES-GCM.
func (e *AESGCMEncryptor) Encrypt(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return append(nonce, ciphertext...), nil
}

// Decrypt decrypts ciphertext using AES-GCM.
func (e *AESGCMEncryptor) Decrypt(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return aesgcm.Open(nil, nonce, ciphertext, nil)
}

// Name returns the name of the algorithm.
func (e *AESGCMEncryptor) Name() string {
	return "AES-GCM"
}

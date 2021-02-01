package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// Encryptor is an interface for struct implementing Encrypt
type Encryptor interface {
	Encrypt(rawData string) (string, error)
}

// AES256Encryptor is a struct implementing Encryptor interface
// that will encrypt provided data using AES256
type AES256Encryptor struct {
	secretKey string
}

// NewAES256Encryptor creates a new instance of AES256Encryptor
func NewAES256Encryptor(secret string) *AES256Encryptor {
	return &AES256Encryptor{
		secretKey: secret,
	}
}

// Encrypt will encrypt a provided raw data into an encrypted string
// using AES256
func (e *AES256Encryptor) Encrypt(rawData string) (string, error) {
	bKey, _ := hex.DecodeString(e.secretKey)
	bRawData := []byte(rawData)

	cipherBlock, err := aes.NewCipher(bKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, bRawData, nil)
	return hex.EncodeToString(cipherText), nil
}

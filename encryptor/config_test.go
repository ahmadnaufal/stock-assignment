package encryptor_test

import (
	"os"
	"testing"

	"github.com/ahmadnaufal/stock-assignment/encryptor"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	os.Setenv("APP_PORT", "8080")
	os.Setenv("AES256_SECRET_KEY", "12345678901234567890123456789012")

	cfg := encryptor.LoadConfig()

	assert.Equal(t, "8080", cfg.Port)
	assert.Equal(t, "12345678901234567890123456789012", cfg.AES256.SecretKey)

	// case if AES256_SECRET_KEY is not defined

	os.Setenv("AES256_SECRET_KEY", "")

	assert.Panics(t, func() { encryptor.LoadConfig() })
}

package encryptor_test

import (
	"testing"

	"github.com/ahmadnaufal/stock-assignment/encryptor"
	"github.com/stretchr/testify/assert"
)

func TestAES256Encryptor_Encrypt(t *testing.T) {
	rawData := "{\"raw_data\":\"test raw data\"}"

	testcases := []struct {
		name    string
		secret  string
		isError bool
	}{
		{
			name:    "success encrypt text",
			secret:  "12345678901234567890123456789012",
			isError: false,
		},
		{
			name:    "invalid case",
			secret:  "11", // 2 bytes key is invalid for AES
			isError: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			enc := encryptor.NewAES256Encryptor(tc.secret)
			cipherText, err := enc.Encrypt(rawData)

			if tc.isError {
				assert.Empty(t, cipherText)
				assert.Error(t, err)
			} else {
				assert.NotEmpty(t, cipherText)
				assert.NoError(t, err)
			}
		})
	}
}

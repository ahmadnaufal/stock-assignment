package encryptor_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahmadnaufal/stock-assignment/encryptor"
	"github.com/ahmadnaufal/stock-assignment/encryptor/mocks"
	"github.com/pkg/errors"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	testcases := []struct {
		name           string
		reqBody        string
		expectedStatus int
		isError        bool
		mockFunction   func(m *mocks.Encryptor)
	}{
		{
			name:           "encrypt success",
			reqBody:        `{"plain_text":"test"}`,
			expectedStatus: http.StatusOK,
			isError:        false,
			mockFunction: func(m *mocks.Encryptor) {
				m.On("Encrypt", "test").Return("example_ciphered_text", nil)
			},
		},
		{
			name:           "no data provided",
			reqBody:        `{"plain_text":""}`,
			expectedStatus: http.StatusBadRequest,
			isError:        true,
		},
		{
			name:           "invalid field body type",
			reqBody:        `{"plain_text":123}`,
			expectedStatus: http.StatusBadRequest,
			isError:        true,
		},
		{
			name:           "error",
			reqBody:        `{"plain_text":"test"}`,
			expectedStatus: http.StatusInternalServerError,
			mockFunction: func(m *mocks.Encryptor) {
				m.On("Encrypt", "test").Return("", errors.New("error while encrypting"))
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockEncryptor := &mocks.Encryptor{}
			if tc.mockFunction != nil {
				tc.mockFunction(mockEncryptor)
			}

			enc := encryptor.NewEncryptorHandler(mockEncryptor)
			req, _ := http.NewRequest("POST", "/encrypt", bytes.NewBuffer([]byte(tc.reqBody)))

			router := httprouter.New()
			enc.Register(router)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code)

			mockEncryptor.AssertExpectations(t)
		})
	}
}

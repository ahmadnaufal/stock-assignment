package encryptor

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

type EncryptorHandler struct {
	encryptor Encryptor
}

type encryptRequest struct {
	PlainText string `json:"plain_text"`
}

type encryptPayload struct {
	CipherText string `json:"cipher_text"`
}

func NewEncryptorHandler(encryptor Encryptor) EncryptorHandler {
	return EncryptorHandler{encryptor: encryptor}
}

func (h *EncryptorHandler) Register(r *httprouter.Router) error {
	if r == nil {
		return errors.New("Router must not be nil")
	}

	r.POST("/encrypt", h.Encrypt)

	return nil
}

func (h *EncryptorHandler) Encrypt(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	reqBody := encryptRequest{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		Error(w, err, http.StatusBadRequest)
		return
	}

	if len(reqBody.PlainText) == 0 {
		Error(w, errors.New("Plain text must not be empty"), http.StatusBadRequest)
		return
	}

	cipherText, err := h.encryptor.Encrypt(reqBody.PlainText)
	if err != nil {
		Error(w, err, http.StatusInternalServerError)
		return
	}

	payload := encryptPayload{
		CipherText: cipherText,
	}

	OK(w, payload)
}

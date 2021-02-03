package stocks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type Encryptor interface {
	PostEncrypt(rawData string) (string, error)
}

type EncryptorService struct {
	host   string
	client *http.Client
}

type encryptRequest struct {
	PlainText string `json:"plain_text"`
}

type encryptResponse struct {
	Status  string `json:"status"`
	Payload struct {
		CipherText string `json:"cipher_text"`
	} `json:"payload"`
}

func NewEncryptorService(host string, client *http.Client) *EncryptorService {
	return &EncryptorService{
		host:   host,
		client: client,
	}
}

func (e *EncryptorService) PostEncrypt(rawData string) (string, error) {
	path := fmt.Sprintf("%s/encrypt", e.host)
	reqBody, _ := json.Marshal(encryptRequest{
		PlainText: rawData,
	})

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}

	resp, err := e.client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusBadRequest {
		errBody, _ := ioutil.ReadAll(resp.Body)
		return "", errors.Wrap(errors.New(string(errBody)), fmt.Sprintf("Error from Encryptor service [%d]", resp.StatusCode))
	}

	var encryptResp encryptResponse
	if err = json.NewDecoder(resp.Body).Decode(&encryptResp); err != nil {
		return "", err
	}

	return encryptResp.Payload.CipherText, nil
}

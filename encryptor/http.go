package encryptor

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status  string      `json:"status"`
	Payload interface{} `json:"payload,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// OK is a wrapper to return 200 OK responses
func OK(w http.ResponseWriter, data interface{}) {
	response := response{
		Status:  "OK",
		Payload: data,
	}
	write(w, response, http.StatusOK)
}

func Error(w http.ResponseWriter, err error, status int) {
	response := response{
		Status:  "ERROR",
		Payload: err.Error(),
	}
	write(w, response, status)
}

func write(w http.ResponseWriter, result interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(result)
}

package stocks

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

type StockHandler struct {
	encryptor Encryptor
	repo      StockRepository
}

type stockPayload struct {
	CipherText string `json:"cipher_text"`
}

func NewStockHandler(encryptor Encryptor, repo StockRepository) StockHandler {
	return StockHandler{
		encryptor: encryptor,
		repo:      repo,
	}
}

func (h *StockHandler) Register(r *httprouter.Router) error {
	if r == nil {
		return errors.New("Router must not be nil")
	}

	r.GET("/stocks/:symbol", h.GetStockSymbol)

	return nil
}

func (h *StockHandler) GetStockSymbol(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	symbol := p.ByName("symbol")
	if symbol == "" {
		Error(w, errors.New("Symbol must be defined"), http.StatusBadRequest)
		return
	}

	stockObj, err := h.repo.FetchBySymbol(symbol)
	if err != nil {
		Error(w, err, http.StatusInternalServerError)
		return
	}

	// encrypt into plain string text to send it to the encryption function
	jsonStr, _ := json.Marshal(stockObj)
	cipherText, err := h.encryptor.PostEncrypt(string(jsonStr))
	if err != nil {
		Error(w, err, http.StatusInternalServerError)
		return
	}

	payload := stockPayload{
		CipherText: cipherText,
	}

	OK(w, payload)
}

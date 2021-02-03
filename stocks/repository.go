package stocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const (
	paramFunction = "TIME_SERIES_DAILY_ADJUSTED"
)

type StockRepository interface {
	FetchBySymbol(symbol string) (StockObject, error)
}

type AlphaVantage struct {
	getStockPath string
	apiKey       string
	client       *http.Client
}

// NewAlphaVantage returns a new instance of AlphaVantage connections
func NewAlphaVantage(getStockPath, apiKey string, client *http.Client) *AlphaVantage {
	return &AlphaVantage{
		getStockPath: getStockPath,
		apiKey:       apiKey,
		client:       client,
	}
}

// FetchBySymbol calls the Alpha Vantage API to retrieve stock data by symbol
func (a *AlphaVantage) FetchBySymbol(symbol string) (StockObject, error) {
	queryParams := a.buildStockQueryParams(symbol)
	path := fmt.Sprintf("%s?%s", a.getStockPath, queryParams)

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return StockObject{}, err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return StockObject{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusBadRequest {
		errBody, _ := ioutil.ReadAll(resp.Body)
		return StockObject{}, errors.Wrap(errors.New(string(errBody)), fmt.Sprintf("Error from AlphaVantage stock service [%d]", resp.StatusCode))
	}

	var responseBody AlphaVantageResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return StockObject{}, err
	}

	return responseBody.ConvertToStockObject(), nil
}

func (a *AlphaVantage) buildStockQueryParams(symbol string) string {
	params := url.Values{}
	params.Add("apikey", a.apiKey)
	params.Add("function", paramFunction)
	params.Add("symbol", symbol)

	return params.Encode()
}

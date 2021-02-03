package stocks

import (
	"strconv"
	"time"
)

type StockObject struct {
	Symbol       string    `json:"symbol"`
	Time         time.Time `json:"time"`
	CurrentValue float64   `json:"current_value"`
	Open         float64   `json:"open"`
	High         float64   `json:"high"`
	Low          float64   `json:"low"`
	Close        float64   `json:"close"`
	Volume       float64   `json:"volume"`
}

type AlphaVantageResponse struct {
	MetaData   AlphaVantageStockMetadata        `json:"Meta Data"`
	TimeSeries map[string]AlphaVantageTimePoint `json:"Time Series (Daily)"`
}

type AlphaVantageStockMetadata struct {
	Symbol   string `json:"2. Symbol"`
	Timezone string `json:"5. Time Zone"`
}

type AlphaVantageTimePoint struct {
	Open             string `json:"1. open"`
	High             string `json:"2. high"`
	Low              string `json:"3. low"`
	Close            string `json:"4. close"`
	AdjustedClose    string `json:"5. adjusted close"`
	Volume           string `json:"6. volume"`
	DividendAmount   string `json:"7. dividend amount"`
	SplitCoefficient string `json:"8. split coefficient"`
}

func (a *AlphaVantageResponse) ConvertToStockObject() (so StockObject) {
	so.Symbol = a.MetaData.Symbol
	for datetime, point := range a.TimeSeries {
		so.Time, _ = time.Parse("2021-02-03", datetime)
		so.Open, _ = strconv.ParseFloat(point.Open, 64)
		so.High, _ = strconv.ParseFloat(point.High, 64)
		so.Low, _ = strconv.ParseFloat(point.Low, 64)
		so.Close, _ = strconv.ParseFloat(point.Close, 64)
		so.Volume, _ = strconv.ParseFloat(point.Volume, 64)

		break
	}

	return
}

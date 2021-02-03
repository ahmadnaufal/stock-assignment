package stocks

import (
	"github.com/joeshaw/envdecode"
	"github.com/subosito/gotenv"
)

type Config struct {
	Port string `env:"APP_PORT,default=8080"`

	Encryptor struct {
		Host    string `env:"ENCRYPTOR_HOST,required"`
		Timeout uint   `env:"ENCRYPTOR_TIMEOUT,default=5"`
	}

	AlphaVantage struct {
		GetStockSymbolURL string `env:"ALPHA_VANTAGE_GET_STOCK_SYMBOL_URL,required"`
		APIKey            string `env:"ALPHA_VANTAGE_API_KEY,required"`
		Timeout           uint   `env:"ALPHA_VANTAGE_TIMEOUT,default=5"`
	}
}

func LoadConfig() Config {
	var config Config
	gotenv.Load()
	if err := envdecode.Decode(&config); err != nil {
		panic(err)
	}

	return config
}

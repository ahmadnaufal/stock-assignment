package encryptor

import (
	"github.com/joeshaw/envdecode"
	"github.com/subosito/gotenv"
)

// Config stores all app-related env vars
type Config struct {
	Port string `env:"APP_PORT,default=8081"`

	AES256 struct {
		SecretKey string `env:"AES256_SECRET_KEY,required"`
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

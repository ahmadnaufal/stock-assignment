package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ahmadnaufal/stock-assignment/stocks"
	"github.com/julienschmidt/httprouter"
)

func main() {
	config := stocks.LoadConfig()

	stockRepo := stocks.NewAlphaVantage(
		config.AlphaVantage.GetStockSymbolURL,
		config.AlphaVantage.APIKey,
		&http.Client{
			Timeout: time.Duration(config.AlphaVantage.Timeout) * time.Second,
		},
	)

	encryptor := stocks.NewEncryptorService(
		config.Encryptor.Host,
		&http.Client{
			Timeout: time.Duration(config.Encryptor.Timeout) * time.Second,
		},
	)

	stockHandler := stocks.NewStockHandler(encryptor, stockRepo)

	router := httprouter.New()
	stockHandler.Register(router)

	s := http.Server{
		Addr:         fmt.Sprintf(":%s", config.Port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func(s *http.Server) {
		log.Printf("Stocks listening at %s\n", s.Addr)
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}(&s)

	<-sigChan

	log.Println("Stocks gracefully stopped.")
}

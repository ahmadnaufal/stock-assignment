package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ahmadnaufal/stock-assignment/encryptor"
	"github.com/julienschmidt/httprouter"
)

func main() {
	config := encryptor.LoadConfig()

	encryptorFlow := encryptor.NewAES256Encryptor(config.AES256.SecretKey)
	encryptorHandler := encryptor.NewEncryptorHandler(encryptorFlow)

	router := httprouter.New()
	encryptorHandler.Register(router)

	s := http.Server{
		Addr:         fmt.Sprintf(":%s", config.Port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func(s *http.Server) {
		log.Printf("Encryptor listening at %s\n", s.Addr)
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}(&s)

	<-sigChan

	log.Println("Encryptor gracefully stopped.")
}

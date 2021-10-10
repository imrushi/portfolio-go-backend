package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/imrushi/portfolio-go-backend/handler"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})

	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
}

func main() {
	log.Info("Initializing Routes...")
	r := mux.NewRouter()

	r.HandleFunc("/v1/health", handler.Health).Methods("GET")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Info("Routes initalized")
	log.Infof("Server is running on port : %s", os.Getenv("API_PORT"))

	s := &http.Server{
		Addr:    fmt.Sprintf(":%v", os.Getenv("API_PORT")),
		Handler: loggedRouter,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("Server failed to start : %s", err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Info("Recieved terminate, graceful shutdown ", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

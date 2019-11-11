package server

import (
	"log"
	"net/http"
	"time"
)

func NewServer(mux *http.ServeMux, serverAddress string, logger *log.Logger) *http.Server {
	srv := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
		ErrorLog:     logger,
	}
	return srv
}

package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"todo-list/config"
	"todo-list/server"
)

func main() {
	var logger = log.New(os.Stdout, "TodoList ", log.Ldate|log.Ltime|log.Llongfile)

	var err = godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}
	var appConfig = config.BuildConfig()

	var mux = http.NewServeMux()

	var serverInstance = server.NewServer(mux, appConfig.Address, logger)

	logger.Print("Server starting")
	err = serverInstance.ListenAndServe()

	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}

	logger.Printf("Server started at: %s", appConfig.Address)
}

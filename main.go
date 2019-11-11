package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"todo-list/config"
)

func main() {
	logger := log.New(os.Stdout, "TodoList ", log.Ldate|log.Ltime|log.Llongfile)

	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}
	var appConfig = config.BuildConfig()

	logger.Print(appConfig.Address)
}

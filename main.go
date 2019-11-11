package main

import (
	"github.com/joho/godotenv"
	"log"
	"todo-list/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}
	config.BuildConfig()
}

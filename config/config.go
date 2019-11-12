package config

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type ServerConfig struct {
	ServerAddress   string
	DatabaseAddress string
	DatabaseClient  *mongo.Database
	DatabaseName    string
}

var AppConfig *ServerConfig = nil

func BuildConfig() *ServerConfig {
	if AppConfig != nil {
		return AppConfig
	}

	AppConfig = &ServerConfig{
		ServerAddress:   os.Getenv("SERVER_ADDRESS"),
		DatabaseAddress: os.Getenv("DATABASE_ADDRESS"),
		DatabaseName:    os.Getenv("DATABASE_NAME"),
	}

	return AppConfig
}

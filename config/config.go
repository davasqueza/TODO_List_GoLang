package config

import "os"

type ServerConfig struct {
	ServerAddress   string
	DatabaseAddress string
}

var AppConfig *ServerConfig = nil

func BuildConfig() *ServerConfig {
	if AppConfig != nil {
		return AppConfig
	}

	AppConfig = &ServerConfig{
		ServerAddress:   os.Getenv("SERVER_ADDRESS"),
		DatabaseAddress: os.Getenv("DATABASE_ADDRESS"),
	}

	return AppConfig
}

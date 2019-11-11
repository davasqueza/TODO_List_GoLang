package config

import "os"

type ServerConfig struct {
	Address string
}

var AppConfig *ServerConfig = nil

func BuildConfig() *ServerConfig {
	if AppConfig != nil {
		return AppConfig
	}

	AppConfig = &ServerConfig{
		Address: os.Getenv("SERVER_ADDRESS"),
	}

	return AppConfig
}

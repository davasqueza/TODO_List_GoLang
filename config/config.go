package config

import "os"

type ServerConfig struct {
	Address string
}

var AppConfig ServerConfig

func BuildConfig() ServerConfig {
	AppConfig.Address = os.Getenv("SERVER_ADDRESS")

	return AppConfig
}

package config

import "github.com/stakkato95/service-engineering-go-lib/config"

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	JwtSecret  string `mapstructure:"JWT_SECRET"`
}

var AppConfig Config

func init() {
	config.Init(&AppConfig, Config{})
}

package config

import "github.com/stakkato95/service-engineering-go-lib/config"

type Config struct {
	ServerPort   string `mapstructure:"SERVER_PORT"`
	UsersService string `mapstructure:"USERS_SERVICE"`
}

var AppConfig Config

func init() {
	config.Init(&AppConfig, Config{})
}

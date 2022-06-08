package config

import "github.com/stakkato95/service-engineering-go-lib/config"

type Config struct {
	ServerPort   string `mapstructure:"SERVER_PORT"`
	DbSource     string `mapstructure:"DB_SOURCE"`
	KafkaTopic   string `mapstructure:"KAFKA_TOPIC"`
	KafkaService string `mapstructure:"KAFKA_SERVICE"`
}

var AppConfig Config

func init() {
	config.Init(&AppConfig, Config{})
}

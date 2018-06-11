package config

import (
	"github.com/spf13/viper"
	"log"
)

func Configure(file string) *Configuration {
	viper.SetConfigName(file)
	viper.AddConfigPath(".")
	var config Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode configuration, %v", err)
	}

	return &config
}

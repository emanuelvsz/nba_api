// config/config.go
package config

import (
	"github.com/spf13/viper"
	"log"
)

func LoadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func GetEnvVar(key string) string {
	return viper.GetString(key)
}

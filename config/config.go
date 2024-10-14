package config

import (
	"log"

	"github.com/spf13/viper"
)

var c Config

func LoadEnv(filename, ext, path string) {
	viper.SetConfigName(filename)
	viper.SetConfigType(ext)
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed read env: %s", err.Error())
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("failed unmarshal env: %s", err.Error())
	}
}

func GetAppPort() string {
	return c.AppPort
}

func GetAppEnv() string {
	return c.AppMode
}

func GetJWTSecret() string {
	return c.JWTSecret
}

func GetPostgresDSN() string {
	return c.Postgres.DSN
}

package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App    AppConfig
	Logger LoggerConfig
	Mongo  MongoConfig
}

type AppConfig struct {
	AppVersion string
	Mode       string
	Port       int
	Debug      bool
}

type LoggerConfig struct {
	Level             string
	Encoding          string
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
}

type MongoConfig struct {
	Uri string
}

func InitConfig(path, filename string) (*Config, error) {
	// Load Config from file
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(path)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("Config file not found")
		}
		return nil, err
	}

	cfg := Config{}
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
		return nil, err
	}

	return &cfg, err
}

package internal

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	LogLevel int    `json:"log_level"`
}

func InitConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Print(err)
		return nil, err
	}

	cfg := &Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Print(err)
		return nil, err
	}

	return cfg, nil
}

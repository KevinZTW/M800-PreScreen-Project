package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

type Config struct {
	MongoDB struct {
		Url  string
		Name string
	}
	LineMessageAPI struct {
		ChannelID          string
		ChannelSecrete     string
		ChannelAccessToken string
	}
}

var (
	once   sync.Once
	config *Config
)

func New() *Config {
	once.Do(func() {
		config = &Config{}

		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./internal/pkg/config")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("fail to read config: ", err.Error())
		}
		if err := viper.Unmarshal(&config); err != nil {
			log.Fatal("fail to decode config: ", err.Error())
		}

	})
	return config
}

package config

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type Config struct {
	Version string `mapstructure:"VERSION"`
	App     App    `mapstructure:",squash"`
	Logger  Logger `mapstructure:",squash"`
}

func NewConfig() *Config {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(home)
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	return config
}

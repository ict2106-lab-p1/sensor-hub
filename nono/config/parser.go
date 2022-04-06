package config

import "github.com/spf13/viper"

func ParseConfig() Config {
	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return config
}

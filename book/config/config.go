package config

import "github.com/spf13/viper"

func getConfig() (*viper.Viper, error) {
	config := viper.New()
	config.SetConfigFile("book/config/config.yaml")
	err := config.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return config, nil
}

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type databaseConfig struct {
	user   string
	dbname string
	host   string
	port   string
	ssl    string
}

func GetConnectionString() (string, error) {
	viper, err := getConfig()
	if err != nil {
		return "", err
	}
	dbConfig := getDatabaseConfig(viper)
	return fmt.Sprintf("user=%s dbname=%s host=%s port=%s sslmode=%s", dbConfig.user, dbConfig.dbname, dbConfig.host, dbConfig.port, dbConfig.ssl), nil
}

func getDatabaseConfig(viper *viper.Viper) *databaseConfig {
	return &databaseConfig{
		user:   viper.GetString("DBUSER"),
		dbname: viper.GetString("DBNAME"),
		host:   viper.GetString("DBHOST"),
		port:   viper.GetString("DBPASS"),
		ssl:    viper.GetString("DBSSL"),
	}
}

package config

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type MySQLConfig struct {
	Name     string
	Host     string
	Port     string
	Username string
	Password string
}

type AppConfig struct {
	Port        string
	SecretKey   string
	MySQLConfig MySQLConfig
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfigs() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig != nil {
		return appConfig
	}

	appConfig = initConfig()

	return appConfig
}

func initConfig() *AppConfig {
	var finalConfig AppConfig
	viper.BindEnv("Port", "APP_PORT")
	viper.BindEnv("SecretKey", "APP_SECRET_KEY")
	viper.BindEnv("MySQLConfig.Name", "MYSQL_DB_NAME")
	viper.BindEnv("MySQLConfig.Host", "MYSQL_DB_HOST")
	viper.BindEnv("MySQLConfig.Port", "MYSQL_DB_PORT")
	viper.BindEnv("MySQLConfig.Username", "MYSQL_DB_USERNAME")
	viper.BindEnv("MySQLConfig.Password", "MYSQL_DB_PASSWORD")
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("failed to extract config, will use default value")
	}
	return &finalConfig
}

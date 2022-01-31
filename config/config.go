package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	RouterEngine   *gin.Engine
	ApiBaseUrl     string
	RunMigration   string
	DataSourceName string
}

func viperEnv(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

func NewConfig() *Config {
	config := new(Config)
	runMigration := viperEnv("DB_MIGRATION")
	apiHost := viperEnv("API_HOST")
	apiPort := viperEnv("API_PORT")

	dbHost := viperEnv("DB_HOST")
	dbPort := viperEnv("DB_PORT")
	dbName := viperEnv("DB_NAME")
	dbUser := viperEnv("DB_USER")
	dbPassword := viperEnv("DB_PASSWORD")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println(dsn)
	config.DataSourceName = dsn

	r := gin.Default()
	config.RouterEngine = r

	config.ApiBaseUrl = fmt.Sprintf("%s:%s", apiHost, apiPort)
	config.RunMigration = runMigration
	return config
}

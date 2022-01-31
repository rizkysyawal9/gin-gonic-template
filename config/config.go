package config

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	RouterEngine   *gin.Engine
	ApiBaseUrl     string
	RunMigration   string
	DataSourceName string
	// TableManagementConfig TableManagementConfig
}

// type TableManagementConfig struct {
// 	ApiBaseUrl string
// }

func NewConfig() *Config {
	config := new(Config)
	runMigration := os.Getenv("DB_MIGRATION")
	apiHost := os.Getenv("API_HOST")
	apiPort := os.Getenv("API_PORT")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	config.DataSourceName = dsn

	// tableManagementBaseUrl := os.Getenv("TABLE_API")
	// tableManagementConfig := TableManagementConfig{ApiBaseUrl: tableManagementBaseUrl}
	// config.TableManagementConfig = tableManagementConfig

	r := gin.Default()
	config.RouterEngine = r

	config.ApiBaseUrl = fmt.Sprintf("%s:%s", apiHost, apiPort)
	config.RunMigration = runMigration
	return config
}

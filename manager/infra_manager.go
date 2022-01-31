package manager

import (
	"log"
	"menu-manage/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Infra interface {
	SqlDb() *gorm.DB
	// HttpClient() *http.Client
	// Config() *config.Config
}

type infra struct {
	db *gorm.DB
	// config *config.Config
}

func NewInfra(config *config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Panicln(err)
	}
	return &infra{
		db: resource,
	}
}

// func (i *infra) HttpClient() *http.Client {
// 	netClient := &http.Client{
// 		Timeout: time.Second * 10,
// 	}
// 	return netClient
// }
// func (i *infra) Config() *config.Config {
// 	return i.config
// }

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	conn, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return conn, nil
}

package api

import (
	"database/sql"
	"log"
	"travelezat-dev/config"
	"travelezat-dev/delivery"
	"travelezat-dev/entity"
	"travelezat-dev/manager"
)

type Server interface {
	Run()
}

type server struct {
	config  *config.Config
	infra   manager.Infra
	usecase manager.UseCaseManager
}

func (s *server) Run() {
	/**
	If we want to run migration just set the DB_MIGRATION env to y or Y
	Note: Gin won't run unless the DB_MIGRATION env is not Y or y
	*/
	if !(s.config.RunMigration == "Y" || s.config.RunMigration == "y") {
		db, _ := s.infra.SqlDb().DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				log.Panic(err)
			}
		}(db)
		s.InitRouter()
		err := s.config.RouterEngine.Run(s.config.ApiBaseUrl)
		if err != nil {
			log.Panic(err)
		}
	} else {
		db := s.infra.SqlDb()
		err := db.AutoMigrate(&entity.Menu{})
		if err != nil {
			log.Panic(err)
		}
		/**
		Input migration dummy data here
		*/
	}
}

func (s *server) InitRouter() {
	publicRoute := s.config.RouterEngine.Group("/api")
	// NewMenuApi(publicRoute, s.usecase.MenuUseCase())
	delivery.NewMenuApi(publicRoute, s.usecase.MenuUseCase())
}

func NewApiServer() Server {
	appConfig := config.NewConfig()
	infra := manager.NewInfra(appConfig)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUseCaseManager(repo)

	return &server{
		config:  appConfig,
		infra:   infra,
		usecase: usecase,
	}
}

package api

import (
	"database/sql"
	"log"
	"menu-manage/config"
	"menu-manage/delivery"
	"menu-manage/entity"
	"menu-manage/manager"
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
	if !(s.config.RunMigration == "Y" || s.config.RunMigration == "y") {
		db, _ := s.infra.SqlDb().DB()
		// err := s.infra.SqlDb().AutoMigrate(&entity.Menu{}, &entity.CustomerTableTransaction{}, &entity.CustomerTable{})
		// if err != nil {
		// 	log.Panic(err)
		// }
		// s.infra.SqlDb().Model(&entity.CustomerTable{}).Save([]entity.CustomerTable{
		// 	{
		// 		ID: "T01",
		// 	},
		// 	{
		// 		ID: "T02",
		// 	},
		// 	{
		// 		ID: "T03",
		// 	},
		// })
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
		err := db.AutoMigrate(&entity.Menu{}, &entity.CustomerTableTransaction{}, &entity.CustomerTable{})
		if err != nil {
			log.Panic(err)
		}
		db.Unscoped().Where("id like ?", "%%").Delete(entity.Menu{})
		db.Model(&entity.CustomerTable{}).Save([]entity.CustomerTable{
			{
				ID: "T01",
			},
			{
				ID: "T02",
			},
			{
				ID: "T03",
			},
		})
		db.Model(&entity.Menu{}).Save([]entity.Menu{
			{
				ID:       "M0001",
				MenuName: "Sayur Kankung",
				Price:    2000,
			},
			{
				ID:       "M0002",
				MenuName: "Sayur Lodeh",
				Price:    3000,
			},
			{
				ID:       "M0003",
				MenuName: "Sayur Jengkol",
				Price:    5000,
			},
		})

	}
}

func (s *server) InitRouter() {
	publicRoute := s.config.RouterEngine.Group("/api")
	// NewMenuApi(publicRoute, s.usecase.MenuUseCase())
	delivery.NewMenuApi(publicRoute, s.usecase.MenuUseCase())
	delivery.NewCustomerTableApi(publicRoute, s.usecase.CustomerTableUseCase())
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

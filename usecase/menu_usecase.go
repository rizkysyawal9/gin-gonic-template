package usecase

import (
	"travelezat-dev/entity"
	"travelezat-dev/repo"
)

type MenuUseCase interface {
	GetAll() ([]entity.Menu, error)
	SearchMenuByName(name string) ([]entity.Menu, error)
	SearchMenuById(id string) (*entity.Menu, error)
}

type menuUseCase struct {
	menuRepo repo.MenuRepository
}

func NewMenuUseCase(menuRepo repo.MenuRepository) MenuUseCase {
	return &menuUseCase{
		menuRepo: menuRepo,
	}
}

func (m *menuUseCase) GetAll() ([]entity.Menu, error) {
	return m.menuRepo.GetAll()
}
func (m *menuUseCase) SearchMenuByName(name string) ([]entity.Menu, error) {
	return m.menuRepo.GetByName(name)
}
func (m *menuUseCase) SearchMenuById(id string) (*entity.Menu, error) {
	return m.menuRepo.GetById(id)
}

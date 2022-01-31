package repo

import (
	"fmt"
	"log"
	"menu-manage/entity"

	"gorm.io/gorm"
)

type MenuRepository interface {
	GetAll() ([]entity.Menu, error)
	GetByName(name string) ([]entity.Menu, error)
	GetById(id string) (*entity.Menu, error)
}

type menuRepository struct {
	db *gorm.DB
}

func (m *menuRepository) GetAll() ([]entity.Menu, error) {
	var list []entity.Menu
	err := m.db.Find(&list).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return list, nil
}
func (m *menuRepository) GetByName(name string) ([]entity.Menu, error) {
	var listByName []entity.Menu
	var searchKeyword = fmt.Sprintf("%%%s%%", name)
	err := m.db.Where("menu_name ilike ?", searchKeyword).
		Find(&listByName).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return listByName, nil
}
func (m *menuRepository) GetById(id string) (*entity.Menu, error) {
	var menu entity.Menu
	err := m.db.Where("id = ?", id).Find(&menu).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &menu, nil
}

func NewMenuRepository(resource *gorm.DB) MenuRepository {
	return &menuRepository{
		db: resource,
	}
}

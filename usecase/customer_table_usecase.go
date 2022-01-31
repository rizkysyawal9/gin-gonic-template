package usecase

import (
	"fmt"
	"menu-manage/apperror"
	"menu-manage/dto"
	"menu-manage/entity"
	"menu-manage/enums"
	"menu-manage/repo"

	"gorm.io/gorm"
)

type CustomerTableUseCase interface {
	GetTodayListCustomerTable() (*[]dto.TableAvailability, error)
	TableCheckIn(checkInRequest dto.Request) (*entity.CustomerTableTransaction, error)
	TableCheckOut(billNo string) error
}

type customerTableUseCase struct {
	repo repo.TableTransactionRepository
}

func NewCustomerTableUseCase(repo repo.TableTransactionRepository) CustomerTableUseCase {
	return &customerTableUseCase{
		repo: repo,
	}
}

func (c *customerTableUseCase) GetTodayListCustomerTable() (*[]dto.TableAvailability, error) {
	return c.repo.GetByBusinessDate()
}
func (c *customerTableUseCase) TableCheckIn(checkInRequest dto.Request) (*entity.CustomerTableTransaction, error) {
	tbl, err := c.repo.CountByTableIdAndStatus(checkInRequest.TableId, enums.TableOccupied)
	if err != nil {
		return nil, apperror.CheckInErr
	}
	if tbl == 0 {
		newTable := &entity.CustomerTableTransaction{
			BillNo:          checkInRequest.BillNo,
			CustomerTableID: checkInRequest.TableId,
			Model:           gorm.Model{},
		}
		table, err := c.repo.CreateOne(newTable)
		if err != nil {
			fmt.Println(err)
			return newTable, err
		}
		return table, nil
	} else {
		return nil, apperror.TableOccupiedErr
	}
	// return
}
func (c *customerTableUseCase) TableCheckOut(billNo string) error {
	return c.repo.Delete(billNo)
}

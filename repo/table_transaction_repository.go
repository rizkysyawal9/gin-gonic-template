package repo

import (
	"errors"
	"fmt"
	"log"
	"menu-manage/apperror"
	"menu-manage/dto"
	"menu-manage/entity"
	"menu-manage/enums"
	"menu-manage/util"

	"gorm.io/gorm"
)

type TableTransactionRepository interface {
	CreateOne(trx *entity.CustomerTableTransaction) (*entity.CustomerTableTransaction, error)
	GetByBusinessDate() (*[]dto.TableAvailability, error)
	CountByTableIdAndStatus(tableId string, tableStatus enums.TableStatus) (int64, error)
	Delete(billno string) error
}

type tableTransactionRepository struct {
	db *gorm.DB
}

func NewTableTransactionRepository(db *gorm.DB) TableTransactionRepository {
	return &tableTransactionRepository{
		db: db,
	}
}

func (t *tableTransactionRepository) CreateOne(trx *entity.CustomerTableTransaction) (*entity.CustomerTableTransaction, error) {

	result := t.db.Find(&trx, "bill_no = ?", trx.BillNo)
	if result.RowsAffected > 0 {
		return trx, errors.New("duplicate Billing Number")
	}
	result = t.db.Find(&entity.CustomerTable{}, "id = ?", trx.CustomerTableID)
	if result.RowsAffected == 0 {
		return trx, errors.New("table doesn't exist")
	}
	err := t.db.Create(&trx).Error
	if err != nil {
		log.Panic(err)
		return trx, err
	}
	fmt.Println(trx)
	return trx, nil
}

func (t *tableTransactionRepository) GetByBusinessDate() (*[]dto.TableAvailability, error) {
	var tableListResult []dto.TableAvailability
	sd, ed := util.GetTodayWithTime()
	err := t.db.Model(&entity.CustomerTable{}).
		Select("customer_table.id as table_id, count(customer_table_transaction.created_at) as is_occupied").
		Group("customer_table.id").
		Joins(`left join customer_table_transaction on customer_table_transaction.customer_table_id  = customer_table.id 
			and customer_table_transaction.created_at between ? and ? 
			and customer_table_transaction.deleted_at is null`,
			sd, ed).
		Scan(&tableListResult).Error
	if err != nil {
		return nil, err
	}
	return &tableListResult, nil
}

func (t *tableTransactionRepository) CountByTableIdAndStatus(tableId string, tableStatus enums.TableStatus) (int64, error) {
	var count int64
	switch tableStatus {
	case enums.TableAllStatus:
		err := t.db.Model(&entity.CustomerTableTransaction{}).Where("customer_table_id = ?", tableId).Count(&count).Error
		if err != nil {
			return -1, err
		}
	case enums.TableVacant:
		err := t.db.Model(&entity.CustomerTableTransaction{}).Where("customer_table_id = ? and deleted_at is not null", tableId).Count(&count).Error
		if err != nil {
			return -1, err
		}
	case enums.TableOccupied:
		err := t.db.Model(&entity.CustomerTableTransaction{}).Where("customer_table_id = ? and deleted_at is null", tableId).Count(&count).Error
		if err != nil {
			return -1, err
		}
	default:
		return -1, apperror.UnknownTableStatusErr
	}
	return count, nil
}
func (t *tableTransactionRepository) Delete(billno string) error {
	result := t.db.Where("bill_no =?", billno).Delete(&entity.CustomerTableTransaction{})
	if result.RowsAffected == 0 {
		return apperror.NoRecordFoundErr
	}
	if result.Error != nil {
		log.Panic(result.Error)
		return result.Error
	}
	return nil
}

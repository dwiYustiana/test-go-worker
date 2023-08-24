package repository

import (
	"test-go-worker/db/entity"

	"github.com/jinzhu/gorm"
)

type TransactionRepoImpl struct {
	DB *gorm.DB
}

type TransactionInterface interface {
	// InsertTransaction(entity.Transaction) error
	InsertTransaction(transaction entity.Transaction) error
}

func NewTransactionRepository(DB *gorm.DB) TransactionInterface {
	return &TransactionRepoImpl{DB}
}

func (e *TransactionRepoImpl) InsertTransaction(transaction entity.Transaction) error {

	tx := e.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(&transaction).Error; err != nil {
		e.DB.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil

}

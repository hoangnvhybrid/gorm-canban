package models

import (
	"gorm-canban/config"
	"gorm-canban/entities"
)

type PointAvailableTransactionRepository interface {
	Create(pointPendingTransaction *entities.PointAvailableTransaction) error
}
type pointAvailableTransactionModel struct {
}

func NewPointAvailableTransactionRepository() PointAvailableTransactionRepository {
	return &pointAvailableTransactionModel{}
}
func (*pointAvailableTransactionModel) Create(pointPendingTransactionModel *entities.PointAvailableTransaction) error {
	db := config.Connection()
	return db.Create(pointPendingTransactionModel).Error
}

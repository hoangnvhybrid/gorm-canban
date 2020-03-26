package models

import (
	"gorm-canban/config"
	"gorm-canban/entities"
)

type PointPendingTransactionRepository interface {
	Create(pointPendingTransaction *entities.PointPendingTransaction) error
}
type pointPendingTransactionModel struct {
}

func NewPointPendingTransactionRepository() PointPendingTransactionRepository {
	return &pointPendingTransactionModel{}
}
func (*pointPendingTransactionModel) Create(pointPendingTransactionModel *entities.PointPendingTransaction) error {
	db := config.Connection()
	return db.Create(pointPendingTransactionModel).Error
}

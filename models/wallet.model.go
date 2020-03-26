package models

import (
	"gorm-canban/config"
	"gorm-canban/entities"
)

type WalletRepository interface {
	Create(wallet *entities.Wallet) error
	FindAll() ([]entities.Wallet, error)
}
type walletModel struct {
}

func NewWalletRepository() WalletRepository {
	return &walletModel{}
}
func (walletModel *walletModel) Create(wallet *entities.Wallet) error {
	db := config.Connection()
	return db.Create(wallet).Error
}
func (walletModel *walletModel) FindAll() ([]entities.Wallet, error) {
	db := config.Connection()
	var wallets []entities.Wallet
	if err := db.Find(&wallets).Error; err != nil {
		return nil, err
	}
	return wallets, nil
}

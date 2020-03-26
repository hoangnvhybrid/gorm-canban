package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm-canban/entities"
	"time"
)

var (
	writeConnection *gorm.DB
	connStr         string
)

// Connection
func Connection() (db *gorm.DB) {
	if writeConnection == nil {
		writeConnection = newConnection()
	}
	db = writeConnection
	return db
}
func newConnection() (db *gorm.DB) {
	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		"127.0.0.1",
		"postgres",
		"fullstack_api",
		"postgres")
	connStr = connectionString
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	db.DB()
	db.DB().SetConnMaxLifetime(time.Hour * 24)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(50)
	db.AutoMigrate(&entities.Wallet{}, &entities.PointAvailableTransaction{}, &entities.PointPendingTransaction{})
	return db
}

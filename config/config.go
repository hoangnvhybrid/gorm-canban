package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDB() (*gorm.DB, error) {
	dbDriver := "mysql"
	dbName := "learngorm"
	dbUser := "root"
	dbPass := ""
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?charset=utf8&parseTime=true")
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return db, nil
}

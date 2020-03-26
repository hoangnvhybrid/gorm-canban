package models

import (
	"gorm-canban/config"
	"gorm-canban/entities"
)

type FacultyModel struct {
}

func (FacultyModel) FindAll() ([]entities.Faculty, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	var faculties []entities.Faculty
	db.Preload("Students").Find(&faculties)
	//db.Find(&faculties)
	return faculties, nil
}

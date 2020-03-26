package entities

import "fmt"

type Faculty struct {
	Id       int `gorm: "primary_key, AUTO_INCREMENT"`
	Name     string
	Students []Student `gorm:"ForeignKey:FacultyId"`
}

func (faculty *Faculty) TableName() string {
	return "faculty"
}

func (faculty Faculty) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s", faculty.Id, faculty.Name)
}

package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Employee struct {
	gorm.Model
	EmpNo     int32     `json:"emp_no" gorm:"emp_no"`
	BirthDate time.Time `json:"birth_date" gorm:"birth_date"`
	FirstName string    `json:"first_name" gorm:"first_name"`
	LastName  string    `json:"last_name" gorm:"last_name"`
	Gender    string    `json:"gender" gorm:"gender"`
}

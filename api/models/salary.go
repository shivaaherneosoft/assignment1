package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type EmployeeSalary struct {
	gorm.Model
	EmpNo    int32 `gorm:"emp_no" json:"emp_no"`
	Salary   int32 `gorm:"salary" json:"salary"`
	FromDate time.Time
	ToDate   time.Time
}

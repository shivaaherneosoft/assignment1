package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type DeptEmp struct {
	gorm.Model
	EmpNo    int32 `gorm:"emp_no" json:"emp_no"`
	DeptNo   int32 `gorm:"dept_no" json:"dept_no"`
	FromDate time.Time
	ToDate   time.Time
}

type DeptManager struct {
	gorm.Model
	EmpNo    int32 `gorm:"emp_no" json:"emp_no"`
	DeptNo   int32 `gorm:"dept_no" json:"dept_no"`
	FromDate time.Time
	ToDate   time.Time
}

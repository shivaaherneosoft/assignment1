package models

import "github.com/jinzhu/gorm"

type Department struct {
	gorm.Model
	DeptName string  `gorm:"dept_name" json:"dept_name"`
	DeptEmp  DeptEmp `json:"-"`
}

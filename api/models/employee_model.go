package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Employee struct {
	gorm.Model
	BirthDate   time.Time `json:"birth_date" gorm:"birth_date"`
	FirstName   string    `json:"first_name" gorm:"first_name"`
	LastName    string    `json:"last_name" gorm:"last_name"`
	Gender      string    `json:"gender" gorm:"gender"`
	HireDate    time.Time `json:"hire_date" gorm:"hire_date"`
	DeptEmp     DeptEmp
	Salary      EmployeeSalary
	DeptManager DeptManager
	Titles      Titles
}

type Titles struct {
	gorm.Model
	EmpNo    int32  `gorm:"emp_no" json:"emp_no"`
	Title    string `gorm:"title" json:"title"`
	FromDate time.Time
	ToDate   time.Time
}

type EmployeeResponse struct {
	EmpID      int32          `json:"id" gorm:"id"`
	BirthDate  time.Time      `json:"birth_date" gorm:"birth_date"`
	FirstName  string         `json:"first_name" gorm:"first_name"`
	LastName   string         `json:"last_name" gorm:"last_name"`
	Gender     string         `json:"gender" gorm:"gender"`
	HireDate   time.Time      `json:"hire_date" gorm:"hire_date"`
	Department Department     `json:"department"`
	Salary     EmployeeSalary `json:"salary"`
}

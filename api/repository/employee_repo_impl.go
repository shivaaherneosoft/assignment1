package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/shivaaherneosoft/assignment1/api/models"
)

type EmployeeRepoIMPL struct {
	Db *gorm.DB
}

func NewEmployeeRepo(db *gorm.DB) EmployeeRepoIMPL {
	return EmployeeRepoIMPL{Db: db}
}

func (e *EmployeeRepoIMPL) Create(emp models.Employee) error {
	return nil
}

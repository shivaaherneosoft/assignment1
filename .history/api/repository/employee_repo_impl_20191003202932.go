package repository

import (
	"fmt"

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
	createUser := e.Db.Table("employees").Create(&emp)
	if createUser.Error != nil {
		fmt.Println("error ", createUser.Error)
	}

	return nil
}

func (e *EmployeeRepoIMPL) Read(empno int32) error {
	createUser := e.Db.Table("employees").Create(&emp)
	if createUser.Error != nil {
		fmt.Println("error ", createUser.Error)
	}

	return nil
}

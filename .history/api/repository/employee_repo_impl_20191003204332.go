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

//Read -
func (e *EmployeeRepoIMPL) Read(empno int32) error {
	emp := models.Employee{}
	getEmployee := e.Db.Table("employees").Where("emp_no = ?", empno).Find(&emp)
	if getEmployee.Error != nil {
		fmt.Println("error ", getEmployee.Error)
		return getEmployee.Error
	}

	return nil
}

//Edit -
func (e *EmployeeRepoIMPL) Edit(emp models.Employee) error {
	e.Db.First(&emp)
	updateEmp := e.Db.Save(&emp)

	if updateEmp.Error != nil {
		fmt.Println("error ", updateEmp.Error)
		return updateEmp.Error
	}

	return nil
}

//Delete -
func (e *EmployeeRepoIMPL) Delete(emp models.Employee) error {
	e.Db.Delete(&emp)

	if updateEmp.Error != nil {
		fmt.Println("error ", updateEmp.Error)
		return updateEmp.Error
	}

	return nil
}

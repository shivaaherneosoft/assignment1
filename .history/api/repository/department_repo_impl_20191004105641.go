package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/shivaaherneosoft/assignment1/api/models"
)

//DepartmentRepoImpl -
type DepartmentRepoImpl struct {
	DB *gorm.DB
}

//NewDepartmentRepo -
func NewDepartmentRepo(db *gorm.DB) DepartmentRepoImpl {
	return DepartmentRepoImpl{DB: db}
}

func (d *DepartmentRepoImpl) Create(dept models.Department) error {
	createDept := d.DB.Table("department").Create(&dept)
	if createDept.Error != nil {
		fmt.Println("[Repo Error] :", createDept.Error)
		return createDept.Error
	}
	return nil
}

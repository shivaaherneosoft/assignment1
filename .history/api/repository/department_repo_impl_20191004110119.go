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

//Create - create new department
func (d *DepartmentRepoImpl) Create(dept models.Department) error {
	createDept := d.DB.Table("department").Create(&dept)
	if createDept.Error != nil {
		fmt.Println("[Repo Error] :", createDept.Error)
		return createDept.Error
	}
	return nil
}

//Read - get specific deparment
func (d *DepartmentRepoImpl) Read(deptno int32) (models.Department, error) {
	dept := models.Department{}
	getDepartment := d.DB.Table("department").Where("dept_no = ?", deptno).Find(&dept)

}

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
	createDept := d.DB.Table("departments").Create(&dept)
	if createDept.Error != nil {
		fmt.Println("[Repo Error] :", createDept.Error)
		return createDept.Error
	}
	return nil
}

//Read - get specific deparment
func (d *DepartmentRepoImpl) Read(deptno int32) (models.Department, error) {
	dept := models.Department{}
	getDepartment := d.DB.Table("departments").Where("id = ?", deptno).Find(&dept)
	if getDepartment.Error != nil {
		fmt.Println("[Repo Error] :", getDepartment.Error)
		return dept, getDepartment.Error
	}

	return dept, nil
}

//Edit - edit depart details
func (d *DepartmentRepoImpl) Update(dept models.Department) error {
<<<<<<< Updated upstream
	updateDept := d.DB.Model(&dept).Omit("id").Updates(&dept)
=======
	// d.DB.First(&dept)
	// updateDept := d.DB.Save(&dept)
	updateDept := d.DB.Model(&dept).Omit("dept_no").Updates(&dept)
>>>>>>> Stashed changes
	if updateDept.Error != nil {
		fmt.Println("error ", updateDept.Error)
		return updateDept.Error
	}
	return nil
}

//Delete -
func (d *DepartmentRepoImpl) Delete(dept models.Department) error {
	deleteDept := d.DB.Delete(&dept)

	if deleteDept.Error != nil {
		fmt.Println("error ", deleteDept.Error)
		return deleteDept.Error
	}

	return nil
}

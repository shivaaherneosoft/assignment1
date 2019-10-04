package repository

import (
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

}

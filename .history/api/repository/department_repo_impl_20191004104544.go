package repository

import (
	"github.com/jinzhu/gorm"
)


type DepartmentRepoImpl struct{
	DB *gorm.DB
}

type NewDepartmentRepo(db *gorm.DB)DepartmentRepoImpl{
	return DepartmentRepoImpl{DB:db}
}
package repository

import "github.com/shivaaherneosoft/assignment1/api/models"

type EmployeeRepo interface {
	Create(models.Employee) error
	Read(int32) (models.Employee, error)
	Update(models.Employee) error
}

type DepartmentRepo interface {
	Create(models.Department) error
	Read(int32) (models.Department, error)
}

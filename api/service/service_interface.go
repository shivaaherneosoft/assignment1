package service

import "github.com/shivaaherneosoft/assignment1/api/models"

type EmployeeService interface {
	Create([]models.Employee) error
	Read(int32) (models.Employee, error)
	Update([]models.Employee) error
}

type DepartmentService interface {
	Create([]models.Department) error
	Read(int32) (models.Department, error)
}

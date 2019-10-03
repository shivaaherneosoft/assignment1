package service

import "github.com/shivaaherneosoft/assignment1/api/models"

type EmployeeService interface {
	Create(models.Employee) error
}

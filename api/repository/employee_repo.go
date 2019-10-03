package repository

import "github.com/shivaaherneosoft/assignment1/api/models"

type EmployeeRepo interface {
	Create(models.Employee) error
}

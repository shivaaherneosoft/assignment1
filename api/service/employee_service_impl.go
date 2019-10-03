package service

import (
	"github.com/shivaaherneosoft/assignment1/api/models"
	"github.com/shivaaherneosoft/assignment1/api/repository"
)

type EmployeeServiceIMPL struct {
	EmployeeRepo repository.EmployeeRepo
}

func NewEmployeeService(repo repository.EmployeeRepo) EmployeeServiceIMPL {
	return EmployeeServiceIMPL{EmployeeRepo: repo}
}

func (e *EmployeeServiceIMPL) Create(employees []models.Employee) error {
	for _, emp := range employees {
		if err := e.EmployeeRepo.Create(emp); err != nil {
			return err
		}
	}
	return nil
}

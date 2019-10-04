package service

import (
	"github.com/shivaaherneosoft/assignment1/api/models"
	"github.com/shivaaherneosoft/assignment1/api/repository"
)

type DepartmentServiceIMPL struct {
	DepartmentRepo repository.DepartmentRepo
}

func NewDepartmentService(repo repository.DepartmentRepo) DepartmentServiceIMPL {
	return DepartmentServiceIMPL{DepartmentRepo: repo}
}

func (e *DepartmentServiceIMPL) Create(departments []models.Department) error {
	for _, dept := range departments {
		if err := e.DepartmentRepo.Create(dept); err != nil {
			return err
		}
	}
	return nil
}

func (e *DepartmentServiceIMPL) Read(empid int32) (models.Department, error) {
	return e.DepartmentRepo.Read(empid)
}

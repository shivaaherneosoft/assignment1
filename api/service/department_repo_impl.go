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

func (d *DepartmentServiceIMPL) Create(departments []models.Department) error {
	for _, dept := range departments {
		if err := d.DepartmentRepo.Create(dept); err != nil {
			return err
		}
	}
	return nil
}

func (d *DepartmentServiceIMPL) Read(empid int32) (models.Department, error) {
	return d.DepartmentRepo.Read(empid)
}

func (d *DepartmentServiceIMPL) Update(departments []models.Department) error {
	for _, dept := range departments {
		if err := d.DepartmentRepo.Update(dept); err != nil {
			return err
		}
	}
	return nil
}

func (d *DepartmentServiceIMPL) Delete(empid int32) error {
	return d.DepartmentRepo.Delete(empid)
}

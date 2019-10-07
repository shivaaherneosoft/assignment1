package repository

import "github.com/shivaaherneosoft/assignment1/api/models"

type EmployeeRepo interface {
	Create(models.Employee) error
	Read(int32) (models.Employee, error)
	Update(models.Employee) error
	Delete(int32) error
}

type DepartmentRepo interface {
<<<<<<< Updated upstream
	Create(models.Department) error
	Read(int32) (models.Department, error)
	Update(models.Department) error
=======
	Create()
	Read(int32) (models.Department, error)
	//	Update()
	//	Delete()
>>>>>>> Stashed changes
}

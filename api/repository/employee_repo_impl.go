package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/shivaaherneosoft/assignment1/api/models"
)

type EmployeeRepoIMPL struct {
	Db *gorm.DB
}

func NewEmployeeRepo(db *gorm.DB) EmployeeRepoIMPL {
	return EmployeeRepoIMPL{Db: db}
}

func (e *EmployeeRepoIMPL) Create(emp models.Employee) error {
	createUser := e.Db.Table("employees").Create(&emp)
	if createUser.Error != nil {
		fmt.Println("error ", createUser.Error)
		return createUser.Error
	}
	emp.DeptEmp.EmpNo = int32(emp.ID)
	deptEmp := e.Db.Table("dept_emps").Create(&emp.DeptEmp)
	if deptEmp.Error != nil {
		fmt.Println("error1 ", deptEmp.Error)
		return deptEmp.Error
	}

	emp.Salary.EmpNo = int32(emp.ID)
	salaryEmp := e.Db.Table("employee_salaries").Create(&emp.Salary)
	if salaryEmp.Error != nil {
		fmt.Println("error2 ", salaryEmp.Error)
		return salaryEmp.Error
	}

	emp.Titles.EmpNo = int32(emp.ID)
	titlesEmp := e.Db.Table("titles").Create(&emp.Titles)
	if titlesEmp.Error != nil {
		fmt.Println("error2 ", titlesEmp.Error)
		return titlesEmp.Error
	}

	return nil
}

//Read -
func (e *EmployeeRepoIMPL) Read(empno int32) (models.EmployeeResponse, error) {
	emp := models.Employee{}
	dept := models.Department{}
	sal := models.EmployeeSalary{}
	response := models.EmployeeResponse{}
	//getEmployee := e.Db.Table("employees").Where("id = ?", empno).Find(&emp)
	getEmployee := e.Db.Table("employees").
		Select("employees.*").
		Joins("left join dept_emps on dept_emps.emp_no = employees.id").
		Joins("left join departments on departments.id = dept_emps.dept_no").
		Where("employees.id = ?", empno).
		Find(&emp)

	if getEmployee.Error != nil {
		fmt.Println("error ", getEmployee.Error)
		return response, getEmployee.Error
	}

	response.EmpID = int32(emp.ID)
	response.BirthDate = emp.BirthDate
	response.FirstName = emp.FirstName
	response.LastName = emp.LastName
	response.Gender = emp.Gender
	response.HireDate = emp.HireDate

	getDepartment := e.Db.Table("employees").
		Select("departments.*").
		Joins("left join dept_emps on dept_emps.emp_no = employees.id").
		Joins("left join departments on departments.id = dept_emps.dept_no").
		Where("employees.id = ?", empno).
		Find(&dept)

	response.Department = dept

	if getDepartment.Error != nil {
		fmt.Println("error ", getDepartment.Error)
		return response, getDepartment.Error
	}

	getSalary := e.Db.Table("employees").
		Select("employee_salaries.*").
		Joins("left join employee_salaries on employee_salaries.emp_no = employees.id").
		Where("employees.id = ?", empno).
		Find(&sal)

	response.Salary = sal

	if getSalary.Error != nil {
		fmt.Println("error ", getSalary.Error)
		return response, getSalary.Error
	}

	return response, nil
}

//Update -
func (e *EmployeeRepoIMPL) Update(emp models.Employee) error {
	// e.Db.First(&emp)
	// updateEmp := e.Db.Save(&emp)
	updateEmp := e.Db.Model(&emp).Omit("id").Updates(&emp)
	if updateEmp.Error != nil {
		fmt.Println("error ", updateEmp.Error)
		return updateEmp.Error
	}
	return nil
}

//Delete -
func (e *EmployeeRepoIMPL) Delete(empid int32) error {
	emp := models.Employee{}
	deleteEmp := e.Db.Where("age = ?", empid).Delete(&emp)
	if deleteEmp.Error != nil {
		fmt.Println("error ", deleteEmp.Error)
		return deleteEmp.Error
	}
	return nil
}

package mysqlpkg

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shivaaherneosoft/assignment1/api/models"
)

func Open() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:root@/organization?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("ERROR:", err)
		return nil, err
	}

	db.AutoMigrate(&models.Employee{}, &models.Department{}, &models.DeptEmp{}, &models.DeptManager{}, &models.Salary{}, &models.Titles{})
	return db, nil
}

package api

import (
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/shivaaherneosoft/assignment1/api/handler"
	"github.com/shivaaherneosoft/assignment1/api/middleware"
	"github.com/shivaaherneosoft/assignment1/api/repository"
	"github.com/shivaaherneosoft/assignment1/api/service"
	"github.com/shivaaherneosoft/assignment1/mysqlpkg"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	db, err := mysqlpkg.Open()
	if err != nil {
		os.Exit(1)
	}

	//Authorization endpoints
	router.POST("/signin", handler.Signin)
	router.POST("/refresh", handler.Refresh)

	//Employee endpoints
	employeeRepo := repository.NewEmployeeRepo(db)
	employeeService := service.NewEmployeeService(&employeeRepo)
	employeeHandler := handler.NewEmployeeHandler(&employeeService)

	router.POST("/E01/employees", middleware.ProtectRequest(employeeHandler.Create))
	router.GET("/E02/employees/:id", middleware.ProtectRequest(employeeHandler.Read))
	router.PATCH("/E03/employees", middleware.ProtectRequest(employeeHandler.Update))
	router.DELETE("/E04/employees", middleware.ProtectRequest(employeeHandler.Update))

	//Department endpoints
	departmentRepo := repository.NewDepartmentRepo(db)
	departmentService := service.NewDepartmentService(&departmentRepo)
	departmentHandler := handler.NewDepartmentHandler(&departmentService)

	router.POST("/E05/departments", middleware.ProtectRequest(departmentHandler.Create))
	router.GET("/E06/departments/:id", middleware.ProtectRequest(departmentHandler.Read))
	router.PATCH("/E07/departments", middleware.ProtectRequest(departmentHandler.Update))

	return router
}

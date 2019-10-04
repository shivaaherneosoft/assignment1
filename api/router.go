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
	repo := repository.NewEmployeeRepo(db)
	service := service.NewEmployeeService(&repo)
	emphandler := handler.NewEmployeeHandler(&service)

	router.POST("/E01/employees", middleware.ProtectRequest(emphandler.Create))
	router.GET("/E02/employees/:id", middleware.ProtectRequest(emphandler.Read))
	router.PATCH("/E03/employees", middleware.ProtectRequest(emphandler.Update))

	router.POST("/signin", handler.Signin)
	router.POST("/refresh", handler.Refresh)
	return router
}

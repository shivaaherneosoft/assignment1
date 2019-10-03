package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shivaaherneosoft/assignment1/api/models"

	"github.com/shivaaherneosoft/assignment1/api/service"

	"github.com/julienschmidt/httprouter"
)

type EmployeeHandler struct {
	EmployeeService service.EmployeeService
}

func NewEmployeeHandler(service service.EmployeeService) EmployeeHandler {
	return EmployeeHandler{EmployeeService: service}
}

func (e *EmployeeHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var employee models.Employee
	err := decoder.Decode(&employee)
	if err != nil {
		panic(err)
	}
	log.Println(employee)

	err = e.EmployeeService.Create(employee)
	fmt.Println("err:", err)
}

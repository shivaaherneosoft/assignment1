package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	var employees []models.Employee
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		fmt.Println("[ERROR]:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid payload"))
		return
	} else {
		err = json.Unmarshal(body, &employees)
		if err != nil {
			fmt.Println("[ERROR]:", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid payload"))
			return
		}

	}

	log.Println("Employee:", employees)

	err = e.EmployeeService.Create(employees)
	if err != nil {
		fmt.Println("[ERROR]:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("employee created successfully!"))
}

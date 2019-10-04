package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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

func (e *EmployeeHandler) Read(w http.ResponseWriter, r *http.Request, path httprouter.Params) {
	pathvariable := path.ByName("id")
	empid, _ := strconv.ParseInt(pathvariable, 10, 32)

	if response, err := e.EmployeeService.Read(int32(empid)); err != nil {
		fmt.Println("[ERROR]:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (e *EmployeeHandler) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

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

	err = e.EmployeeService.Update(employees)
	if err != nil {
		fmt.Println("[ERROR]:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("employee updated successfully!"))
}

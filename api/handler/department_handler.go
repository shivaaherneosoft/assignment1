package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/shivaaherneosoft/assignment1/api/models"
	"github.com/shivaaherneosoft/assignment1/api/service"
)

type DepartmentHandler struct {
	DepartmentService service.DepartmentService
}

func NewDepartmentHandler(service service.DepartmentService) DepartmentHandler {
	return DepartmentHandler{DepartmentService: service}
}

func (d *DepartmentHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var departments []models.Department
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		fmt.Println("[ERROR]:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		err = json.Unmarshal(body, &departments)
		if err != nil {
			fmt.Println("[ERROR]:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	}

	err = d.DepartmentService.Create(departments)
	if err != nil {
		fmt.Println("[ERROR]:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

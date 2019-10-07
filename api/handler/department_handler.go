package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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

func (d *DepartmentHandler) Read(w http.ResponseWriter, r *http.Request, path httprouter.Params) {
	pathvariable := path.ByName("id")
	empid, _ := strconv.ParseInt(pathvariable, 10, 32)

	if response, err := d.DepartmentService.Read(int32(empid)); err != nil {
		fmt.Println("[ERROR]:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (d *DepartmentHandler) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	err = d.DepartmentService.Update(departments)
	if err != nil {
		fmt.Println("[ERROR]:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (d *DepartmentHandler) Delete(w http.ResponseWriter, r *http.Request, path httprouter.Params) {
	pathvariable := path.ByName("id")
	deptno, _ := strconv.ParseInt(pathvariable, 10, 32)

	if err := d.DepartmentService.Delete(int32(deptno)); err != nil {
		fmt.Println("[ERROR]:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

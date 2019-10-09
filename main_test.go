package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/shivaaherneosoft/assignment1/api/handler"
	"github.com/shivaaherneosoft/assignment1/api/repository"
	"github.com/shivaaherneosoft/assignment1/api/service"
	"github.com/shivaaherneosoft/assignment1/mysqlpkg"

	"github.com/gorilla/mux"
)

var a mux.Router

func TestSigninTable(t *testing.T) {

	router := httprouter.New()

	router.POST("/signin", handler.Signin)
	payload := []byte(`{"username":"user1","password":"password1"}`)
	req, err := http.NewRequest("POST", "/signin", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreateEmployeeTable(t *testing.T) {
	router := httprouter.New()
	db, err := mysqlpkg.Open()
	if err != nil {
		t.Errorf("connection failed ")
		os.Exit(1)
	}
	employeeRepo := repository.NewEmployeeRepo(db)
	employeeService := service.NewEmployeeService(&employeeRepo)
	employeeHandler := handler.NewEmployeeHandler(&employeeService)
	router.POST("/employees", employeeHandler.Create)

	payload := []byte(`[{
		"birth_date":"2012-04-23T18:25:43.511Z",
		"first_name":"Rahul K",
		"last_name":"Patil",
		"gender":"male",
		"hire_date":"2012-04-23T18:25:43.511Z",
		"DeptEmp":{"dept_no":1,"FromDate":"2012-04-23T00:00:00.00Z", "ToDate":"2012-04-23T00:00:00.00Z"},
		"Salary":{"salary":333000,"FromDate":"2012-04-23T00:00:00.00Z", "ToDate":"2012-04-23T00:00:00.00Z"},
		"Titles":{"title":"SE","FromDate":"2012-04-23T00:00:00.00Z", "ToDate":"2012-04-23T00:00:00.00Z"}
	}]
	`)
	req, err := http.NewRequest("POST", "/employees", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Add("Authorization", "")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNTcwNjE4Njk1LCJqdGkiOiJhZG1pbiJ9.Lw7wifH6CDbR9kJls7BP6dduDODswT5NJlWANMwNQoM")

	rr := httptest.NewRecorder()
	Token := req.Header.Get("Authorization")

	router.ServeHTTP(rr, req)

	if accessToken := Token; accessToken == "" {
		t.Errorf("Unauthorised token %v", http.StatusUnauthorized)
	}
	fmt.Println("code ", rr.Code)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code and provide valid access token: got %v want %v",
			status, http.StatusCreated)
	}

}

package main

import (
	"log"
	"net/http"

	"github.com/shivaaherneosoft/assignment/api"
)

func main() {
	router := api.NewRouter()
	log.Println(http.ListenAndServe(":5000", router))
}

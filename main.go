package main

import (
	"log"
	"net/http"

	"github.com/shivaaherneosoft/assignment1/api"
	"github.com/shivaaherneosoft/assignment1/config"
)

func init() {
	config.SetConfig()
}

func main() {
	router := api.NewRouter()
	log.Println("starting server...")
	log.Println(http.ListenAndServe(":5000", router))
}

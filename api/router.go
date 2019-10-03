package api

import (
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	// router.GET("/employees", )
	return router
}

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func router() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/", homePageHandler)
	router.HandlerFunc(http.MethodGet, "/getCountryInfo/:country", getCountryInfo)
	return router
}

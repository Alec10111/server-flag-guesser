package routes

import (
	"net/http"
	"fgapi/src/handlers"
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/", handlers.HomePageHandler)
	router.HandlerFunc(http.MethodGet, "/getCountryInfo/:country", handlers.GetCountryInfo)
	return router
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func homePageHandler(w http.ResponseWriter, req *http.Request) {

	jsonRes, err := json.Marshal("Available")
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}
func getCountryInfo(w http.ResponseWriter, req *http.Request) {
	params := httprouter.ParamsFromContext(req.Context())
	info := fetchDescription(params.ByName("country"))

	jsonRes, err := json.Marshal(info)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

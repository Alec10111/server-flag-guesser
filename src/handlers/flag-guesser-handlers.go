package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func enableCors(w *http.ResponseWriter) {
	// (*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}

func HomePageHandler(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	jsonRes, err := json.Marshal("Available")
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}
func GetCountryInfo(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	params := httprouter.ParamsFromContext(req.Context())
	info := FetchDescription(params.ByName("country"))

	jsonRes, err := json.Marshal(info)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

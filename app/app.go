package main

import (
	"github.com/annguyen1892/goservice/app/config"
	"github.com/annguyen1892/goservice/app/handler"
	"github.com/annguyen1892/goservice/app/models"
	"github.com/gorilla/mux"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
)

var config = Config{}
var handler = Handler{}

func resResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()

	handler.Server = config.server
	handler.Database = config.database
	handler.Connect()
}

func AllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := handler.findAll()
	if err != nill {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resResponse(w, http.StatusOK, products)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", AllProducts).Methods("GET")
}

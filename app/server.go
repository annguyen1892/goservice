package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	// . "talaria-recentlyviewed-go/app/config"
	"fmt"
	// "github.com/getsentry/raven-go"
	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent"
	"os"
	. "talaria-recentlyviewed-go/app/handler"
	. "talaria-recentlyviewed-go/app/models"
)

// var conf = Config{}
var handler = Handler{}

func response(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondWithError
func respondWithError(w http.ResponseWriter, code int, msg string) {
	response(w, code, map[string]string{"error": msg})
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	handler.Addrs = os.Getenv("MONGODB_HOST")
	timeout := 30
	timeout, err = strconv.Atoi(os.Getenv("MONGODB_TIMEOUT"))
	if err != nil {
		log.Panic(err)
	}
	handler.Timeout = timeout
	handler.Database = os.Getenv("MONGODB_DATABASE")
	handler.Username = os.Getenv("MONGODB_USERNAME")
	handler.Password = os.Getenv("MONGODB_PASSWORD")
	handler.Replicasetname = os.Getenv("MONGODB_REPLICATESET")
	handler.Readpreference = os.Getenv("MONGODB_READPREFERENCE")
	handler.Connect()

	// Sentry
	// raven.SetDSN(os.Getenv("SENTRY_KEY"))
}

// AllProducts ...
func GetProducts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()

	limit := 20
	if len(query.Get("limit")) > 0 {
		param_limit, err := strconv.Atoi(query.Get("limit"))
		if err == nil {
			limit = param_limit
		}
	}

	page := 1
	if len(query.Get("page")) > 0 {
		param_page, err := strconv.Atoi(query.Get("page"))
		if err == nil {
			page = param_page
		}
	}

	user_id := 0
	if len(params["id"]) > 0 {
		param_uid, err := strconv.Atoi(params["id"])
		if err == nil {
			user_id = param_uid
		}
	}

	products, err := handler.GetList(user_id, limit, page)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response(w, http.StatusOK, products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	product_ids := r.FormValue("product_ids")

	if len(product_ids) == 0 {
		response(w, http.StatusOK, Create{Result: "success"})
		return
	}

	user_id := 0
	if len(params["id"]) > 0 {
		param_uid, err := strconv.Atoi(params["id"])
		if err == nil {
			user_id = param_uid
		}
	}

	products := handler.CreateProducts(user_id, product_ids)

	response(w, http.StatusOK, products)
}

func main() {

	config := newrelic.NewConfig(os.Getenv("APP_NAME"), os.Getenv("NEWRELIC_KEY"))
	app, err := newrelic.NewApplication(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := mux.NewRouter()
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/users/{id}/products", GetProducts)).Methods("GET")
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/users/{id}/products", CreateProduct)).Methods("POST")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

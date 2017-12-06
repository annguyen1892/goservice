package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	. "talaria-recentlyviewed-go/app/models"
	"testing"
)

func TestGetProducts(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/0/products", nil)
	if err != nil {
		log.Panic(err)
		return
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProducts)
	handler.ServeHTTP(rr, req)
	body := rr.Body.String()
	if !strings.Contains(body, "paging") || !strings.Contains(body, "data") {
		t.Errorf("handler returned unexpected body:%v", body)
	}
}

func getResponse(body []byte) (*Response, error) {
	var s = new(Response)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

func TestCreate(t *testing.T) {
	values := url.Values{}
	values.Set("product_ids", "111,222")
	req, err := http.NewRequest("POST", "/users/0/products", strings.NewReader(values.Encode()))
	if err != nil {
		log.Panic(err)
		return
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProduct)
	handler.ServeHTTP(rr, req)
	body := rr.Body.String()
	expected := `{"result":"success"}`
	if expected != body {
		t.Errorf("handler returned unexpected body: expectec %v and got %v", expected, body)
	}
}

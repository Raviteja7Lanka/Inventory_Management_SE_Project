package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestRootEndPoint(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndPoint).Methods("GET")
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	//fmt.Println(response.Body)
	//assert.Equal(t, "Hello World", response.Body.String(), "Hello world message is expected")
}

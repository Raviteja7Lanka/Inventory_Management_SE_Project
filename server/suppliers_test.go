package main

import (
	"apis/controllers"
	"apis/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetSuppliers(t *testing.T) {
	router := mux.NewRouter()
	var suppliers []models.Suppliers

	router.HandleFunc("/supplier/all", controllers.GetAllSuppliers).Methods("GET")
	request, _ := http.NewRequest("GET", "/supplier/all", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &suppliers)

	if err != nil {
		fmt.Println("err is ", err)
	}

	if len(suppliers) >= 0 {
		name := suppliers[0].NAME
		address := suppliers[0].ADDRESS
		fmt.Println(name, address)
		assert.True(t, true, "Atleast one customer order was returned")
	} else {
		assert.Fail(t, "Atleast one order was expected")
	}
}

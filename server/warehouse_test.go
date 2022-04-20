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

func TestGetWarehouses(t *testing.T) {
	router := mux.NewRouter()
	var warehouses []models.Warehouses

	router.HandleFunc("/warehouse/all", controllers.GetAllWarehouses).Methods("GET")
	request, _ := http.NewRequest("GET", "/warehouse/all", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &warehouses)

	if err != nil {
		fmt.Println("err is ", err)
	}

	if len(warehouses) >= 1 {
		name := warehouses[1].WAREHOUSE_ID
		//location cannot be null name quantity
		address := warehouses[1].LOCATION
		fmt.Println(name, address)
		fmt.Println("Hurray!")
		assert.True(t, true, "Atleast one warehouse was returned")
	} else {
		assert.Fail(t, "Atleast one warehouse was expected")
	}
}

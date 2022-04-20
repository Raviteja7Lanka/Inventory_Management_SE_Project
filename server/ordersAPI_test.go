package main

import (
	"apis/controllers"
	"apis/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllCustomerOrders(t *testing.T) {
	router := mux.NewRouter()
	var orders []models.Orders

	router.HandleFunc("/customer/orders/all", controllers.GetAllCustomerOrders).Methods("GET")
	request, _ := http.NewRequest("GET", "/customer/orders/all", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &orders)

	if err != nil {
		fmt.Println("err is ", err)
	}

	if len(orders) >= 1 {
		name := orders[1].OrderId
		address := orders[1].SupplierId
		fmt.Println(name, address)
		assert.True(t, true, "Atleast one customer order was returned")
	} else {
		assert.Fail(t, "Atleast one order was expected")
	}
}

func TestGetCustomerOrdersByID(t *testing.T) {
	router := mux.NewRouter()
	var orders []models.Orders

	router.HandleFunc("/customer/orders/{ordId}", controllers.GetAllCustomerOrders).Methods("GET")
	request, _ := http.NewRequest("GET", "/customer/orders/12", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &orders)

	if err != nil {
		fmt.Println("err is ", err)
	}

	if len(orders) == 1 {
		name := orders[1].OrderId
		address := orders[1].SupplierId
		fmt.Println(name, address)
		assert.True(t, true, "Only one customer order was returned as expected")
	} else {
		assert.Fail(t, "Atleast one order was expected")
	}
}

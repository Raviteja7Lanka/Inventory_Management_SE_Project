package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)


func TestGetAllCustomerOrders(t *testing.T) {
	router := mux.NewRouter()
	// var orders[] Orders

	router.HandleFunc("/customer/orders/all", GetAllCustomerOrders).Methods("GET")
	// request, _ := http.NewRequest("GET", "http://localhost:8085/customer/orders/all", nil)
	// response := httptest.NewRecorder()
	// router.ServeHTTP(response, request)

	
	// 

	//
	// if(response!=nil)
	// {
	// 	assert.NotNil(t,response.Body,"Records found")
	// }
	request, _ := http.NewRequest("GET", "/customer/orders/all", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	fmt.Println(response.Code)
	fmt.Println(response.Body)
	// err := json.Unmarshal([]byte(response.Body.Bytes()), &orders)
	// if err != nil {
	// 		fmt.Println("err is ", err)
	// 	}
		
	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.NotNil(t, response.Body, "Response is not null")
	//fmt.Println(response.Body)
	
}
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

func TestGetOrderDetails(t *testing.T) {
	router := mux.NewRouter()
	var order_details []models.OrderDetails

	router.HandleFunc("/orders/details/all", controllers.GetAllOrderDetails).Methods("GET")
	request, _ := http.NewRequest("GET", "/orders/details/all", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &order_details)

	if err != nil {
		fmt.Println("err is ", err)
	}

	if len(order_details) >= 1 {
		// name := order_details[0]
		// address := order_details[0].ADDRESS
		// fmt.Println(name, address)
		fmt.Println("Atleast one order details was returned")
		assert.True(t, true, "Atleast one order details was returned")
	} else {
		assert.Fail(t, "Atleast one order details was expected")
	}
}

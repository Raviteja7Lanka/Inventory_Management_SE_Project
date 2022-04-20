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

func TestGetPayments(t *testing.T) {
	router := mux.NewRouter()
	var payments []models.Payments

	router.HandleFunc("/payment/all", controllers.GetAllPayments).Methods("GET")
	request, _ := http.NewRequest("GET", "/payment/all", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &payments)

	if err != nil {
		fmt.Println("err is ", err)
	}

	if len(payments) > 0 {
		number := payments[0].BILL_NUMBER
		types := payments[0].PAYMENT_TYPE
		fmt.Println(number, types)
		assert.True(t, true, "Atleast one Payment was returned")
	} else {
		assert.Fail(t, "Atleast one Payment was expected")
	}
}

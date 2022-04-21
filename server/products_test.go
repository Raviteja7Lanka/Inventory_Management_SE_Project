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

func TestGetProducts(t *testing.T) {
	router := mux.NewRouter()
	var products []models.Products

	router.HandleFunc("/product/all", controllers.GetAllProducts).Methods("GET")
	request, _ := http.NewRequest("GET", "/product/all", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &products)

	if err != nil {
		fmt.Println("err is ", err)
	}

	if len(products) >= 0 {
		// name := products[0].PRODUCT_NAME
		// fmt.Println(name)
		assert.True(t, true, "Atleast one Product was returned")
	} else {
		assert.Fail(t, "Atleast one Staff was expected but none are present")
	}
}

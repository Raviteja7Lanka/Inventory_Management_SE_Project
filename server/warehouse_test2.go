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

// func TestGetWarehouseID(t *testing.T) {
// 	router := mux.NewRouter()
// 	var warehouses []models.Warehouses

// 	router.HandleFunc("/warehouse/all", controllers.GetAllWarehouses).Methods("GET")
// 	request, _ := http.NewRequest("GET", "/warehouse/all", nil)
// 	response := httptest.NewRecorder()
// 	router.ServeHTTP(response, request)

// 	//testing for just status code here
// 	assert.Equal(t, 200, response.Code, "OK response is expected")
// 	err := json.Unmarshal([]byte(response.Body.Bytes()), &warehouses)

// 	if err != nil {
// 		fmt.Println("err is ", err)
// 	}
// 	//fmt.Println(societies[1].SocietyName)
// 	if len(warehouses) >= 1 {
// 		name := warehouses[0].WAREHOUSE_ID
// 		address := warehouses[0].LOCATION

// 		//fmt.Println(startDateNotNull, endDateNotNull)
// 		assert.Equal(t, len(name) > 0, true, "Non empty name expected")
// 		assert.Equal(t, len(address) > 0, true, "Non empty address expected")
// 	}
// }

func TestGetWarehousebyID(t *testing.T) {
	router := mux.NewRouter()
	var warehouses models.Warehouses

	router.HandleFunc("/warehouse/{wareId}", controllers.GetWarehouseByID).Methods("GET")
	request, _ := http.NewRequest("GET", "/warehouse/1", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	//testing for just status code here
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &warehouses)

	fmt.Println("response is: ", response.Body.String())
	if err != nil {
		fmt.Println("err is ", err)
	}
	fmt.Println("Warehouse_test2")
	if (warehouses != models.Warehouses{}) {
		fmt.Println("Warehouse_test2")
		assert.Equal(t, 1, int(warehouses.ID), "ID's should be matched")
	}
}

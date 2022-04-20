// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"testing"

// 	"github.com/gorilla/mux"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetCustomerOrder(t *testing.T) {
// 	router := mux.NewRouter()
// 	if router == nil {
// 		fmt.Println("Null Router")
// 	}
// 	var order Orders
// 	router.HandleFunc("/customer/order/14", getCustomerOrderByID).Methods("GET")
// 	resp, err := http.Get("http://localhost:8085/customer/order/14")
// 	if err != nil {
// 		fmt.Println("NULL response")
// 	}
// 	body, er := ioutil.ReadAll(resp.Body)
// 	if er != nil {
// 		e := json.Unmarshal([]byte(string(body)), &order)

// 		if e != nil {
// 			fmt.Println("err is ", e)
// 		}

// 		assert.NotNil(t, order, "Get Order by ID sucsess")
// 	}
// }

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

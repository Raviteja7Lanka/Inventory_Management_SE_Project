package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetCustomerOrder(t *testing.T) {
	router := mux.NewRouter()
	if router == nil {
		fmt.Println("Null Router")
	}
	var order Orders
	router.HandleFunc("/customer/order/14", getCustomerOrderByID).Methods("GET")
	resp, err := http.Get("http://localhost:8085/customer/order/14")
	if err != nil {
		fmt.Println("NULL response")
	}
	body, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		e := json.Unmarshal([]byte(string(body)), &order)

		if e != nil {
			fmt.Println("err is ", e)
		}

		assert.NotNil(t, order, "Get Order by ID sucsess")
	}
}

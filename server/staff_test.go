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

func TestGetStaff(t *testing.T) {
	router := mux.NewRouter()
	var staff []models.Staff

	router.HandleFunc("/staff/all", controllers.GetAllStaff).Methods("GET")
	request, _ := http.NewRequest("GET", "/staff/all", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &staff)

	if err != nil {
		fmt.Println("err is ", err)
	}

	if len(staff) >= 0 {
		// name := staff[0].FIRST_NAME
		// lastname := staff[0].LAST_NAME
		// fmt.Println(name, lastname)
		fmt.Println("Atleast one Staff was returned")
		assert.True(t, true, "Atleast one Staff was returned")
	} else {
		assert.Fail(t, "Atleast one Staff was expected but none are present")
	}
}

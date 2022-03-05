package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// import "gorm.io/gorm"

func RootEndPoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("OK SUCCESS"))
}

func InitRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/customer/orders/all", getAllCustomerOrders).Methods("GET")
	router.HandleFunc("/customer/order/{ordId}", getCustomerOrderByID).Methods("GET")
	router.HandleFunc("/customer/orders/add", addCustomerOrder).Methods("POST")
	router.HandleFunc("/customer/orders/{ordId}", updateCustomerOrder).Methods("PUT")
	router.HandleFunc("/customer/orders/{ordId}", deleteCustomerOrder).Methods("DELETE")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func main() {
	InitDB()
	InitRouter()

}

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// import "gorm.io/gorm"

func RootEndPoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("Hello World"))
	// response.Header().Set("Access-Control-Allow-Origin", "*")
}

func InitRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/customer/orders/all", getAllCustomerOrders).Methods("GET")

	//
	router.HandleFunc("/customer/all", getAllCustomers).Methods("GET")
	router.HandleFunc("/supplier/all", getAllSuppliers).Methods("GET")
	router.HandleFunc("/staff/all", getAllStaff).Methods("GET")
	router.HandleFunc("/product/all", getAllProduct).Methods("GET")
	router.HandleFunc("/payment/all", getAllPayments).Methods("GET")
	//

	router.HandleFunc("/customer/order/{ordId}", getCustomerOrderByID).Methods("GET")

	//
	router.HandleFunc("/customer/{cusId}", getCustomerByID).Methods("GET")
	router.HandleFunc("/supplier/{supId}", getSupplierByID).Methods("GET")
	router.HandleFunc("/staff/{staffId}", getStaffByID).Methods("GET")
	router.HandleFunc("/product/{prodId}", getProductByID).Methods("GET")
	router.HandleFunc("/payment/{payId}", getPaymentsByID).Methods("GET")
	//

	router.HandleFunc("/customer/orders/add", addCustomerOrder).Methods("POST")

	//
	router.HandleFunc("/customer/add", addCustomer).Methods("POST")
	router.HandleFunc("/supplier/add", addSupplier).Methods("POST")
	router.HandleFunc("/staff/add", addStaff).Methods("POST")
	router.HandleFunc("/product/add", addProduct).Methods("POST")
	router.HandleFunc("/payment/add", addPayments).Methods("POST")
	//

	router.HandleFunc("/customer/orders/{ordId}", updateCustomerOrder).Methods("PUT")

	//
	router.HandleFunc("/customer/{cusId}", updateCustomer).Methods("PUT")
	router.HandleFunc("/supplier/{supId}", updateSupplier).Methods("PUT")
	router.HandleFunc("/staff/{staffId}", updateStaff).Methods("PUT")
	router.HandleFunc("/product/{prodId}", updateProduct).Methods("PUT")
	router.HandleFunc("/payment/{payId}", updatePayments).Methods("PUT")
	//

	router.HandleFunc("/customer/orders/{ordId}", deleteCustomerOrder).Methods("DELETE")

	router.HandleFunc("/customer/all", getAllCustomers).Methods("GET")
	router.HandleFunc("/customer/{custId}", getCustomerByID).Methods("GET")
	router.HandleFunc("/customer/add", addCustomer).Methods("POST")
	router.HandleFunc("/customer/{custId}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customer/{custId}", deleteCustomer).Methods("DELETE")

	router.HandleFunc("/supplier/all", getAllSuppliers).Methods("GET")
	router.HandleFunc("/supplier/{supId}", getSupplierByID).Methods("GET")
	router.HandleFunc("/supplier/add", addSupplier).Methods("POST")
	router.HandleFunc("/supplier/{supId}", updateSupplier).Methods("PUT")
	router.HandleFunc("/supplier/{supId}", deleteSupplier).Methods("DELETE")

	// router.HandleFunc("/staff/all", getAllCustomerOrders).Methods("GET")
	router.HandleFunc("/staff/{staffId}", getStaffByID).Methods("GET")
	router.HandleFunc("/staff/add", addStaff).Methods("POST")
	router.HandleFunc("/staff/{ordId}", updateStaff).Methods("PUT")
	router.HandleFunc("/staff/{ordId}", deleteStaff).Methods("DELETE")

	//
	router.HandleFunc("/customer/{cusId}", deleteCustomer).Methods("DELETE")
	router.HandleFunc("/supplier/{supId}", deleteSupplier).Methods("DELETE")
	router.HandleFunc("/staff/{staffId}", deleteStaff).Methods("DELETE")
	router.HandleFunc("/product/{prodId}", deleteProduct).Methods("DELETE")
	router.HandleFunc("/payment/{payId}", deletePayments).Methods("DELETE")
	//

	corsObj := handlers.AllowedOrigins([]string{"http://localhost:4200"})
	headersOk := handlers.AllowedHeaders([]string{"accept", "origin", "X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "PUT", "POST", "DELETE", "OPTIONS", "PATCH"})
	allowCreds := handlers.AllowCredentials()
	allowOptions := handlers.OptionStatusCode(204)

	router.HandleFunc("/", RootEndPoint).Methods("GET")

	http.Handle("/", router)
	// log.Fatal(http.ListenAndServe(":8085", nil))
	log.Fatal(http.ListenAndServe(":8085", handlers.CORS(corsObj, headersOk, methodsOk, allowCreds, allowOptions)(router)))
}

func main() {
	InitDB()
	InitRouter()

}
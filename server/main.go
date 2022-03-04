package main

import "github.com/gorilla/mux"
import "net/http"
import "log"
// import "gorm.io/gorm"


func InitRouter(){
	router:= mux.NewRouter()
	router.HandleFunc("/customer/orders/all", getAllCustomerOrders).Methods("GET")
    router.HandleFunc("/customer/order/{ordId}", getCustomerOrderByID).Methods("GET")
    router.HandleFunc("/customer/orders/add", addCustomerOrder).Methods("POST")
    router.HandleFunc("/customer/orders/{ordId}", updateCustomerOrder).Methods("PUT")
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
	
	
	
	http.Handle("/",router)
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func main() {
	InitDB()
    InitRouter()
	
}
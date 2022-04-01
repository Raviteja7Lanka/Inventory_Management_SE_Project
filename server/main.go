package main

import (
	"log"
	// "apis/models"
	"net/http"

	"apis/routes"

	"github.com/gorilla/mux"
)

// import "gorm.io/gorm"

func RootEndPoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("Hello World"))
}

func InitRouter() {
	router := mux.NewRouter()
	// cust:=&models.Customers{}
	// router.HandleFunc("/customer/orders/all", getAllCustomerOrders).Methods("GET")
	// router.HandleFunc("/customer/order/{ordId}", getCustomerOrderByID).Methods("GET")
	// router.HandleFunc("/customer/orders/{ordId}", updateCustomerOrder).Methods("PUT")
	// router.HandleFunc("/customer/orders/add", addCustomerOrder).Methods("POST")
	// router.HandleFunc("/customer/orders/{ordId}", deleteCustomerOrder).Methods("DELETE")
	routes.RegisterCustomerRoutes(router)

	router.HandleFunc("/supplier/all", getAllSuppliers).Methods("GET")
	router.HandleFunc("/supplier/{supId}", getSupplierByID).Methods("GET")
	router.HandleFunc("/supplier/add", addSupplier).Methods("POST")
	router.HandleFunc("/supplier/{supId}", updateSupplier).Methods("PUT")
	router.HandleFunc("/supplier/{supId}", deleteSupplier).Methods("DELETE")

	router.HandleFunc("/staff/all", getAllStaff).Methods("GET")
	router.HandleFunc("/staff/{staffId}", getStaffByID).Methods("GET")
	router.HandleFunc("/staff/add", addStaff).Methods("POST")
	router.HandleFunc("/staff/{staffId}", updateStaff).Methods("PUT")
	router.HandleFunc("/staff/{staffId}", deleteStaff).Methods("DELETE")

	// router.HandleFunc("/product/all", getAllProduct).Methods("GET")
	// router.HandleFunc("/product/{prodId}", getProductByID).Methods("GET")
	// router.HandleFunc("/product/add", addProduct).Methods("POST")
	// router.HandleFunc("/product/{prodId}", updateProduct).Methods("PUT")
	// router.HandleFunc("/product/{prodId}", deleteProduct).Methods("DELETE")
	routes.RegisterProductRoutes(router)

	//router.HandleFunc("/payment/all", getAllPayments).Methods("GET")
	//router.HandleFunc("/payment/{payId}", getPaymentsByID).Methods("GET")
	//router.HandleFunc("/payment/add", addPayments).Methods("POST")
	//router.HandleFunc("/payment/{payId}", updatePayments).Methods("PUT")
	//router.HandleFunc("/payment/{payId}", deletePayments).Methods("DELETE")
	routes.RegisterPaymentRoutes(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func main() {
	InitDB()
	InitRouter()

}

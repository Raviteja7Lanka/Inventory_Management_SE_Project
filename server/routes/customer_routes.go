package routes

import (
	"apis/controllers"

	"github.com/gorilla/mux"
)

func RegisterCustomerRoutes(router *mux.Router) {

	router.HandleFunc("/customer/all", controllers.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customer/{cusId}", controllers.GetCustomerByID).Methods("GET")
	router.HandleFunc("/customer/add", controllers.AddCustomer).Methods("POST")
	router.HandleFunc("/customer/{cusId}", controllers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customer/{cusId}", controllers.DeleteCustomer).Methods("DELETE")
}

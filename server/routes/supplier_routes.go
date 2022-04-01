package routes

import (
	"apis/controllers"

	"github.com/gorilla/mux"
)

func RegisterSupplierRoutes(router *mux.Router) {

	router.HandleFunc("/supplier/all", controllers.GetAllStaff).Methods("GET")
	router.HandleFunc("/supplier/{supId}", controllers.GetPaymentByID).Methods("GET")
	router.HandleFunc("/supplier/{supId}", controllers.UpdatePayment).Methods("PUT")
	router.HandleFunc("/supplier/add", controllers.AddPayment).Methods("POST")
	router.HandleFunc("/supplier/{supId}", controllers.DeletePayment).Methods("DELETE")
}

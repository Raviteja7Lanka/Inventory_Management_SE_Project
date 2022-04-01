package routes

import (
	"apis/controllers"

	"github.com/gorilla/mux"
)

func RegisterStaffRoutes(router *mux.Router) {

	router.HandleFunc("/staff/all", controllers.GetAllStaff).Methods("GET")
	router.HandleFunc("/staff/{staffId}", controllers.GetPaymentByID).Methods("GET")
	router.HandleFunc("/staff/{staffId}", controllers.UpdatePayment).Methods("PUT")
	router.HandleFunc("/staff/add", controllers.AddPayment).Methods("POST")
	router.HandleFunc("/staff/{staffId}", controllers.DeletePayment).Methods("DELETE")
}

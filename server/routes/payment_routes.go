package routes

import (
	"apis/controllers"

	"github.com/gorilla/mux"
)

func RegisterPaymentRoutes(router *mux.Router) {

	router.HandleFunc("/payment/all", controllers.GetAllPayments).Methods("GET")
	router.HandleFunc("/payment/{payId}", controllers.GetPaymentsByID).Methods("GET")
	router.HandleFunc("/payment/{payId}", controllers.UpdatePayments).Methods("PUT")
	router.HandleFunc("/payment/add", controllers.AddPayments).Methods("POST")
	router.HandleFunc("/payment/{payId}", controllers.DeletePayments).Methods("DELETE")
}

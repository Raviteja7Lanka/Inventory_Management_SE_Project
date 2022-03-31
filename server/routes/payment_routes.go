package routes

import (
	"apis/controllers"

	"github.com/gorilla/mux"
)

func RegisterPaymentRoutes(router *mux.Router) {

	router.HandleFunc("/payment/all", controllers.GetAllPayments).Methods("GET")
	router.HandleFunc("/payment/{payId}", controllers.GetPaymentByID).Methods("GET")
	router.HandleFunc("/payment/add", controllers.UpdatePayment).Methods("PUT")
	router.HandleFunc("/payment/{payId}", controllers.AddPayment).Methods("POST")
	router.HandleFunc("/payment/{payId}", controllers.DeletePayment).Methods("DELETE")
}

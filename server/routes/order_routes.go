package routes

import (
	"apis/controllers"

	"github.com/gorilla/mux"
)

func RegisterCustomerOrderRoutes(router *mux.Router) {

	router.HandleFunc("/customer/orders/all", controllers.GetAllOrders).Methods("GET")
	router.HandleFunc("/customer/order/{ordId}", controllers.GetOrderByID).Methods("GET")
	router.HandleFunc("/customer/orders/{ordId}", controllers.UpdateOrder).Methods("PUT")
	router.HandleFunc("/customer/orders/add", controllers.AddOrder).Methods("POST")
	router.HandleFunc("/customer/orders/{ordId}", controllers.DeleteOrder).Methods("DELETE")
}

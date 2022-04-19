package routes

import (
	"apis/controllers"

	"github.com/gorilla/mux"
)

func RegisterCustomerOrderRoutes(router *mux.Router) {

	router.HandleFunc("/customer/orders/all", controllers.GetAllCustomerOrders).Methods("GET")
	router.HandleFunc("/customer/order/{ordId}", controllers.GetCustomerOrderByID).Methods("GET")
	router.HandleFunc("/customer/order/status/{ordStatus}", controllers.GetAllCustomerOrdersByStatus).Methods("GET")
	router.HandleFunc("/customer/orders/{ordId}", controllers.UpdateCustomerOrder).Methods("PUT")
	router.HandleFunc("/customer/orders/add", controllers.AddCustomerOrder).Methods("POST")
	router.HandleFunc("/customer/orders/{ordId}", controllers.DeleteCustomerOrder).Methods("DELETE")
}

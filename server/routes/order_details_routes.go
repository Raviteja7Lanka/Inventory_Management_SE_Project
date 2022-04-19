package routes

import (
	"apis/controllers"

	"github.com/gorilla/mux"
)

func RegisterOrderDetailsRoutes(router *mux.Router) {

	router.HandleFunc("/orders/details/all", controllers.GetAllOrderDetails).Methods("GET")
	router.HandleFunc("/order/details/{ordId}", controllers.GetOrderDetailsByID).Methods("GET")
	router.HandleFunc("/orders/details/{ordId}", controllers.UpdateOrderDetails).Methods("PUT")
	router.HandleFunc("/orders/details/add", controllers.AddOrderDetails).Methods("POST")
	router.HandleFunc("/orders/details/{ordId}", controllers.DeleteOrderDetails).Methods("DELETE")
}

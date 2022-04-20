package routes

import (
	"apis/controllers"

	"github.com/gorilla/mux"
)

func RegisterOrderRoutes(router *mux.Router) {

	router.HandleFunc("/orders/all", controllers.GetAllOrders).Methods("GET")
	router.HandleFunc("/order/{ordId}", controllers.GetOrderByID).Methods("GET")
	router.HandleFunc("/order/status/{ordStatus}", controllers.GetAllOrdersByStatus).Methods("GET")
	router.HandleFunc("/order/type/{ordType}", controllers.GetAllOrdersByType).Methods("GET")
	router.HandleFunc("/orders/{ordId}", controllers.UpdateOrder).Methods("PUT")
	router.HandleFunc("/orders/add", controllers.AddOrder).Methods("POST")
	router.HandleFunc("/orders/{ordId}", controllers.DeleteOrder).Methods("DELETE")
}

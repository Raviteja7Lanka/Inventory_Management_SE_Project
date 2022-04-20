package routes

import (
	"apis/controllers"

	"github.com/gorilla/mux"
)

func RegisterWarehouseRoutes(router *mux.Router) {

	router.HandleFunc("/warehouse/all", controllers.GetAllWarehouses).Methods("GET")
	router.HandleFunc("/warehouse/{wareId}", controllers.GetWarehouseByID).Methods("GET")
	router.HandleFunc("/warehouse/{wareId}", controllers.UpdateWarehouse).Methods("PUT")
	router.HandleFunc("/warehouse/add", controllers.AddWarehouse).Methods("POST")
	router.HandleFunc("/warehouse/{wareId}", controllers.DeleteWarehouse).Methods("DELETE")
}

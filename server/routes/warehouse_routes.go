package routes

import (
	"github.com/gorilla/mux"
)

func RegisterWarehouseRoutes(router *mux.Router) {

	router.HandleFunc("/warehouse/all", GetAllWarehouses).Methods("GET")
	router.HandleFunc("/warehouse/{wareId}", GetWarehouseByID).Methods("GET")
	router.HandleFunc("/warehouse/add", AddWarehouse).Methods("POST")
	router.HandleFunc("/warehouse/{wareId}", UpdateWarehouse).Methods("PUT")
	router.HandleFunc("/warehouse/{wareId}", DeleteWarehouse).Methods("DELETE")
}

package routes

import (
	"apis/controllers"

	"github.com/gorilla/mux"
)

func RegisterSupplierRoutes(router *mux.Router) {

	router.HandleFunc("/supplier/all", controllers.GetAllSuppliers).Methods("GET")
	router.HandleFunc("/supplier/{supId}", controllers.GetSupplierByID).Methods("GET")
	router.HandleFunc("/supplier/{supId}", controllers.UpdateSupplier).Methods("PUT")
	router.HandleFunc("/supplier/add", controllers.AddSupplier).Methods("POST")
	router.HandleFunc("/supplier/{supId}", controllers.DeleteSupplier).Methods("DELETE")
}

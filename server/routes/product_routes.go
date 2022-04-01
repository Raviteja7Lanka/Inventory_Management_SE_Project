package routes

import (
	"apis/controllers"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {

	router.HandleFunc("/product/all", controllers.GetAllProducts).Methods("GET")
	router.HandleFunc("/product/{prodId}", controllers.GetProductByID).Methods("GET")
	router.HandleFunc("/product/{prodId}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/add", controllers.AddProduct).Methods("POST")
	router.HandleFunc("/product/{prodId}", controllers.DeleteProduct).Methods("DELETE")
}

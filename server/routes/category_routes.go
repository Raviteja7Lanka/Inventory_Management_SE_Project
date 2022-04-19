package routes

import (
	"apis/controllers"

	"github.com/gorilla/mux"
)

func RegisterCategoryRoutes(router *mux.Router) {

	router.HandleFunc("/category/all", controllers.GetAllCategories).Methods("GET")
	router.HandleFunc("/category/{catId}", controllers.GetCategoryByID).Methods("GET")
	router.HandleFunc("/warehouse-categories/{wareId}", controllers.GetCategoriesByWarehouseID).Methods("GET")
	router.HandleFunc("/category/{catId}", controllers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/category/add", controllers.AddCategory).Methods("POST")
	router.HandleFunc("/category/{catId}", controllers.DeleteCategory).Methods("DELETE")
}

package routes

import (
	"apis/controllers"

	"github.com/gorilla/mux"
)

func RegisterCategoryRoutes(router *mux.Router) {

	router.HandleFunc("/category/all", controllers.GetAllCategories).Methods("GET")
	router.HandleFunc("/category/{catID}", controllers.GetCategoriesByID).Methods("GET")
	router.HandleFunc("/category/{catID}", controllers.UpdateCategories).Methods("PUT")
	router.HandleFunc("/category/add", controllers.AddCategories).Methods("POST")
	router.HandleFunc("/category/{catID}", controllers.DeleteCategories).Methods("DELETE")
}

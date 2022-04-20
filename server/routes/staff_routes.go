package routes

import (
	"apis/controllers"

	"github.com/gorilla/mux"
)

func RegisterStaffRoutes(router *mux.Router) {

	router.HandleFunc("/staff/all", controllers.GetAllStaff).Methods("GET")
	router.HandleFunc("/staff/{staffId}", controllers.GetStaffByID).Methods("GET")
	router.HandleFunc("/staff/{staffId}", controllers.UpdateStaff).Methods("PUT")
	router.HandleFunc("/staff/add", controllers.AddStaff).Methods("POST")
	router.HandleFunc("/staff/{staffId}", controllers.DeleteStaff).Methods("DELETE")
}

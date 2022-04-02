package routes

import (
	"apis/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoleRoutes(router *mux.Router) {

	router.HandleFunc("/role/all", controllers.GetAllRoles).Methods("GET")
	router.HandleFunc("/role/{roleID}", controllers.GetRoleByID).Methods("GET")
	router.HandleFunc("/role/{roleID}", controllers.UpdateRole).Methods("PUT")
	router.HandleFunc("/role/add", controllers.AddRole).Methods("POST")
	router.HandleFunc("/role/{roleID}", controllers.DeleteRole).Methods("DELETE")
}

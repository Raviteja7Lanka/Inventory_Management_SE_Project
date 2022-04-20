package controllers

import (
	"apis/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllWarehouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var warehouses []models.Warehouses
	e := db.Find(&warehouses).Error
	if e != nil {
		// sendErr(w, http.StatusInternalServerError, e.Error())
		return
	}
	e = json.NewEncoder(w).Encode(warehouses)
	if e != nil {
		// sendErr(w, http.StatusInternalServerError, e.Error())
	}
}

func GetWarehouseByID(w http.ResponseWriter, r *http.Request) {
	var warehouses models.Warehouses
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["wareId"])
	db.First(&warehouses, queryParams["wareId"])
	e := json.NewEncoder(w).Encode(warehouses)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
	}
}

func AddWarehouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var warehouse models.Warehouses
	json.NewDecoder(r.Body).Decode(&warehouse)
	db.Create(&warehouse)
	json.NewEncoder(w).Encode(warehouse)
}

func UpdateWarehouse(w http.ResponseWriter, r *http.Request) {
	var warehouse models.Warehouses
	var updatedWarehouse models.Warehouses
	queryParams := mux.Vars(r)
	cusId := queryParams["wareId"]
	db.First(&warehouse, cusId)
	json.NewDecoder(r.Body).Decode(&updatedWarehouse)
	db.Model(&warehouse).Where("warehouse_id=?", cusId).Updates(&updatedWarehouse)
	json.NewEncoder(w).Encode(&warehouse)

}

func DeleteWarehouse(w http.ResponseWriter, r *http.Request) {
	// var warehouse models.Warehouses
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["wareId"])
	er := db.Exec("delete from warehouses where warehouse_id=?", queryParams["wareId"]).Error
	if er != nil {
		json.NewEncoder(w).Encode("{Status:201, Message: Internal Error}")
	}
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

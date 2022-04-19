package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Warehouses struct {
	// gorm.Model
	WAREHOUSE_ID string `json:"warehouse_id"`
	NAME         string `json:"name"`
	LOCATION     string `json:"location"`
	DESCRIPTION  string `json:"description"`
	STATUS       string `json:"status"`
}

func getAllWarehouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var warehouses []Warehouses
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

func getWarehouseByID(w http.ResponseWriter, r *http.Request) {
	var warehouses Warehouses
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["wareId"])
	db.First(&warehouses, queryParams["wareId"])
	e := json.NewEncoder(w).Encode(warehouses)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func addWarehouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var warehouse Warehouses
	json.NewDecoder(r.Body).Decode(&warehouse)
	db.Create(&warehouse)
	json.NewEncoder(w).Encode(warehouse)
}

func updateWarehouse(w http.ResponseWriter, r *http.Request) {
	var warehouse Warehouses
	var updatedSupplier Warehouses
	queryParams := mux.Vars(r)
	cusId := queryParams["wareId"]
	db.First(&warehouse, cusId)
	json.NewDecoder(r.Body).Decode(&updatedSupplier)
	db.Model(&warehouse).Where("warehouse_id=?", cusId).Updates(&updatedSupplier)
	json.NewEncoder(w).Encode(&warehouse)

}

func deleteWarehouse(w http.ResponseWriter, r *http.Request) {
	var warehouse Warehouses
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["wareId"])
	db.Delete(&warehouse, queryParams["wareId"])
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

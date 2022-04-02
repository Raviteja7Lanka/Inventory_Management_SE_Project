package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Warehouses struct {
	// gorm.Model
	warehouse_id uint   `json:"warehouse_id"`
	name         string `json:"name"`
	location     string `json:"location"`
	description  string `json:"descriptiom"`
	status       string `json:"status"`
}

func getAllWarehouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var warehouses []Warehouses
	e := db.Find(&warehouses).Error
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	e = json.NewEncoder(w).Encode(warehouses)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
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
	var supplier Warehouses
	json.NewDecoder(r.Body).Decode(&supplier)
	db.Create(&supplier)
	json.NewEncoder(w).Encode(supplier)
}

func updateWarehouse(w http.ResponseWriter, r *http.Request) {
	var supplier Warehouses
	var updatedSupplier Warehouses
	queryParams := mux.Vars(r)
	cusId := queryParams["wareId"]
	db.First(&supplier, cusId)
	json.NewDecoder(r.Body).Decode(&updatedSupplier)
	db.Model(&supplier).Where("customer_id=?", cusId).Updates(&updatedSupplier)
	json.NewEncoder(w).Encode(&supplier)

}

func deleteWarehouse(w http.ResponseWriter, r *http.Request) {
	var supplier Warehouses
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["wareId"])
	db.Delete(&supplier, queryParams["wareId"])
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

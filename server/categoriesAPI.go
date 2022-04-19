package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Categories struct {
	// gorm.Model
	CATEGORY_ID  string `json:"category_id"`
	WAREHOUSE_ID string `json:"category_id"`
	NAME         string `json:"name"`
	DESCRIPTION  string `json:"description"`
}

func getAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var categories []Categories
	e := db.Find(&categories).Error
	if e != nil {
		// sendErr(w, http.StatusInternalServerError, e.Error())
		return
	}
	e = json.NewEncoder(w).Encode(categories)
	if e != nil {
		// sendErr(w, http.StatusInternalServerError, e.Error())
	}
}

func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	var categories Categories
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["wareId"])
	db.First(&categories, queryParams["wareId"])
	e := json.NewEncoder(w).Encode(categories)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func addCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var category Categories
	json.NewDecoder(r.Body).Decode(&category)
	db.Create(&category)
	json.NewEncoder(w).Encode(category)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	var category Categories
	var updatedSupplier Categories
	queryParams := mux.Vars(r)
	cusId := queryParams["wareId"]
	db.First(&category, cusId)
	json.NewDecoder(r.Body).Decode(&updatedSupplier)
	db.Model(&category).Where("category_id=?", cusId).Updates(&updatedSupplier)
	json.NewEncoder(w).Encode(&category)

}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	var category Categories
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["wareId"])
	db.Delete(&category, queryParams["wareId"])
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

package controllers

import (
	"apis/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var categories []models.Categories
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

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	var categories models.Categories
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["catId"])
	db.First(&categories, queryParams["catId"])
	e := json.NewEncoder(w).Encode(categories)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
	}
}

func GetCategoriesByWarehouseID(w http.ResponseWriter, r *http.Request) {
	var categories []models.Categories
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["wareId"])
	db.Where("warehouse_id=?", queryParams["wareId"]).Find(&categories)
	e := json.NewEncoder(w).Encode(categories)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
	}
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var category models.Categories
	json.NewDecoder(r.Body).Decode(&category)
	db.Create(&category)
	json.NewEncoder(w).Encode(category)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Categories
	var updatedCategory models.Categories
	queryParams := mux.Vars(r)
	cusId := queryParams["catId"]
	db.First(&category, cusId)
	json.NewDecoder(r.Body).Decode(&updatedCategory)
	db.Model(&category).Where("category_id=?", cusId).Updates(&updatedCategory)
	json.NewEncoder(w).Encode(&category)

}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// var category models.Categories
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["catId"])
	er := db.Exec("delete from categories where category_id=?", queryParams["catId"]).Error
	if er != nil {
		json.NewEncoder(w).Encode("{Status:201, Message: Internal Error}")
	}
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

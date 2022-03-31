package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Products struct {
	// gorm.Model
	PRODUCTS_ID           string `json:"products_id"`
	PRODUCT_NAME          string `json:"product_name"`
	PRODUCT_DESCR_IP_TION string `json:"product_descr_ip_tion"`
	PRODUCT_UNIT          uint   `json:"product_unit"`
	PRODUCT_PRICE         uint   `json:"product_price"`
	PRODUCT_QUANTITY      uint   `json:"product_quantity"`
	PRODUCT_STATUS        string `json:"product_status"`
	OTHER_DETAILS         string `json:"other_details"`
	SUPPLIER_ID           string `json:"supplier_id"`
	CATEGORY_ID           string `json:"category_id"`
	//

}

func getAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []Products
	e := db.Find(&products).Error
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	e = json.NewEncoder(w).Encode(products)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func getProductByID(w http.ResponseWriter, r *http.Request) {
	var products Products

	queryParams := mux.Vars(r)
	fmt.Println(queryParams["prodId"])
	db.First(&products, queryParams["prodId"])
	e := json.NewEncoder(w).Encode(products)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products Products
	json.NewDecoder(r.Body).Decode(&products)
	db.Create(&products)
	json.NewEncoder(w).Encode(products)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	var products Products
	var updatedProducts Products
	queryParams := mux.Vars(r)
	prodId := queryParams["prodId"]
	db.First(&products, prodId)
	json.NewDecoder(r.Body).Decode(&updatedProducts)
	db.Model(&products).Where("products_id=?", prodId).Updates(&updatedProducts)
	json.NewEncoder(w).Encode(&products)

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	var products Products
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["prodId"])
	db.Delete(&products, queryParams["prodId"])
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

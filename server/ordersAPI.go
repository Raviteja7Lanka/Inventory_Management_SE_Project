package main

import (
	"fmt"
	//     "log"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Orders struct {
	// gorm.Model
	OrderId        uint   `json:"order_id"`
	DateOfOrder    string `json:"date_of_order"`
	OrderDetailsId string `json:"order_details_id"`
	CustomerId     uint   `json:"customer_id"`
	SupplierId     uint   `json:"supplier_id"`
	Status         string `json:"status"`
}

var db *gorm.DB
var err error

func InitDB() {
	db, err = gorm.Open(sqlite.Open("INVENTORY_MANAGEMENT.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Database connection failed")
	}
	//db.AutoMigrate(&Orders{})
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

func getAllCustomerOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	var orders []Orders
	e := db.Find(&orders).Error
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	e = json.NewEncoder(w).Encode(orders)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}

}

func getCustomerOrderByID(w http.ResponseWriter, r *http.Request) {
	var order Orders

	queryParams := mux.Vars(r)
	fmt.Println(queryParams["ordId"])
	db.First(&order, queryParams["ordId"])
	e := json.NewEncoder(w).Encode(order)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}
func addCustomerOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Orders
	json.NewDecoder(r.Body).Decode(&order)
	db.Create(&order)
	json.NewEncoder(w).Encode(order)
	// fmt.Println("Hello")
}
func updateCustomerOrder(w http.ResponseWriter, r *http.Request) {
	var order Orders
	var updatedOrder Orders
	queryParams := mux.Vars(r)
	ordId := queryParams["ordId"]
	db.First(&order, ordId)
	json.NewDecoder(r.Body).Decode(&updatedOrder)
	db.Model(&order).Where("order_id=?", ordId).Updates(&updatedOrder)
	json.NewEncoder(w).Encode(&order)

}

func deleteCustomerOrder(w http.ResponseWriter, r *http.Request) {
	var order Orders
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["ordId"])
	db.Delete(&order, queryParams["ordId"])
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

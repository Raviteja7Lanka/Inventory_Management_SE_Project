package controllers

import (
	"apis/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllOrderDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order_details []models.OrderDetails
	e := db.Find(&order_details).Error
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
		return
	}
	e = json.NewEncoder(w).Encode(order_details)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
	}

}

func GetOrderDetailsByID(w http.ResponseWriter, r *http.Request) {
	var order_details models.OrderDetails
	// db.Raw("select id, date_of_order, ordId, customer_id, supplier_id, status from Orders where id=?",1).Scan(&order)
	// data,_:= json.Marshal(&order)
	// fmt.Fprint(w, string(data))
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["ordId"])
	db.First(&order_details, queryParams["ordId"])
	e := json.NewEncoder(w).Encode(order_details)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
	}
}
func AddOrderDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order_details models.OrderDetails
	json.NewDecoder(r.Body).Decode(&order_details)
	db.Create(&order_details)
	json.NewEncoder(w).Encode(order_details)
	// fmt.Println("Hello")
}
func UpdateOrderDetails(w http.ResponseWriter, r *http.Request) {
	var order_details models.OrderDetails
	var updatedOrderDetails models.OrderDetails
	queryParams := mux.Vars(r)
	ordId := queryParams["ordId"]
	db.First(&order_details, ordId)
	json.NewDecoder(r.Body).Decode(&updatedOrderDetails)
	db.Model(&order_details).Where("order_details_id=?", ordId).Updates(&updatedOrderDetails)
	json.NewEncoder(w).Encode(&order_details)

}

func DeleteOrderDetails(w http.ResponseWriter, r *http.Request) {
	var order_details models.OrderDetails
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["ordId"])
	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
	db.Delete(&order_details, queryParams["ordId"])
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

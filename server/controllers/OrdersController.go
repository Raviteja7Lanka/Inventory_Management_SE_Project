package controllers

import (
	"apis/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// func GetAllOrders(w http.ResponseWriter, r *http.Request) {
// 	// w.Header().Set("Content-Type", "application/json")
// 	Orders := models.GetAllOrders()
// 	res, e := json.Marshal(Orders)
// 	if e != nil {
// 		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
// 		http.Error(w, string(resp), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func GetOrderByID(w http.ResponseWriter, r *http.Request) {
// 	// db.Raw("select id, date_of_order, order_details_id, Order_id, supplier_id, status from models.Orders where id=?",1).Scan(&order)
// 	// data,_:= json.Marshal(&order)
// 	// fmt.Fprint(w, string(data))
// 	queryParams := mux.Vars(r)
// 	// fmt.Println(queryParams["ordId"])
// 	// db.First(&Orders, queryParams["ordId"])
// 	Orders := models.GetOrderByID(queryParams["ordId"])
// 	res, e := json.Marshal(Orders)
// 	if e != nil {
// 		http.Error(w, string(e.Error()), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func AddOrder(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	Orders := &models.Orders{}
// 	json.NewDecoder(r.Body).Decode(&Orders)
// 	newCust := models.AddOrder(Orders)
// 	res, e := json.Marshal(newCust)
// 	if e != nil {
// 		http.Error(w, string(e.Error()), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// 	// fmt.Println("Hello")
// }

// func UpdateOrder(w http.ResponseWriter, r *http.Request) {

// 	updatedOrder := &models.Orders{}
// 	queryParams := mux.Vars(r)
// 	ordId := queryParams["ordId"]
// 	json.NewDecoder(r.Body).Decode(&updatedOrder)
// 	orders := models.UpdateOrder(ordId, updatedOrder)
// 	res, e := json.Marshal(orders)
// 	if e != nil {
// 		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
// 		http.Error(w, string(resp), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }

// func DeleteOrder(w http.ResponseWriter, r *http.Request) {

// 	queryParams := mux.Vars(r)
// 	Orders := models.DeleteOrder(queryParams["cusId"])
// 	//db.Raw("delete from models.Orders where order_id=?", queryParams["ordId"])
// 	res, e := json.Marshal(Orders)
// 	if e != nil {
// 		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
// 		http.Error(w, string(resp), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }

func GetAllCustomerOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	var orders []models.Orders
	e := db.Find(&orders).Error
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
		return
	}
	e = json.NewEncoder(w).Encode(orders)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
	}

}

func GetCustomerOrderByID(w http.ResponseWriter, r *http.Request) {
	var order models.Orders

	queryParams := mux.Vars(r)
	fmt.Println(queryParams["ordId"])
	db.First(&order, queryParams["ordId"])
	e := json.NewEncoder(w).Encode(order)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
	}
}
func AddCustomerOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order models.Orders
	json.NewDecoder(r.Body).Decode(&order)
	db.Create(&order)
	json.NewEncoder(w).Encode(order)
	// fmt.Println("Hello")
}
func UpdateCustomerOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Orders
	var updatedOrder models.Orders
	queryParams := mux.Vars(r)
	ordId := queryParams["ordId"]
	db.First(&order, ordId)
	json.NewDecoder(r.Body).Decode(&updatedOrder)
	db.Model(&order).Where("order_id=?", ordId).Updates(&updatedOrder)
	json.NewEncoder(w).Encode(&order)

}

func DeleteCustomerOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Orders
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["ordId"])
	db.Delete(&order, queryParams["ordId"])
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

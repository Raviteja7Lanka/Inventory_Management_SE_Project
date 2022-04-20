package controllers

import (
	"apis/config"
	"apis/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// func init() {
// 	config.ConnectSqlite()
// 	db = config.GetDB()
// }

// func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
// 	// w.Header().Set("Content-Type", "application/json")
// 	customers := models.GetAllCustomers()
// 	res, e := json.Marshal(customers)
// 	if e != nil {
// 		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
// 		http.Error(w, string(resp), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
// 	// db.Raw("select id, date_of_order, order_details_id, customer_id, supplier_id, status from Orders where id=?",1).Scan(&order)
// 	// data,_:= json.Marshal(&order)
// 	// fmt.Fprint(w, string(data))
// 	queryParams := mux.Vars(r)
// 	// fmt.Println(queryParams["cusId"])
// 	// db.First(&customers, queryParams["cusId"])
// 	customers := models.GetCustomerByID(queryParams["cusId"])
// 	res, e := json.Marshal(customers)
// 	if e != nil {
// 		http.Error(w, string(e.Error()), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func AddCustomer(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	customers := &models.Customers{}
// 	json.NewDecoder(r.Body).Decode(&customers)
// 	newCust := models.AddCustomer(customers)
// 	res, e := json.Marshal(newCust)
// 	if e != nil {
// 		http.Error(w, string(e.Error()), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// 	// fmt.Println("Hello")
// }

// func UpdateCustomer(w http.ResponseWriter, r *http.Request) {

// 	updatedCustomer := &models.Customers{}
// 	queryParams := mux.Vars(r)
// 	cusId := queryParams["cusId"]
// 	json.NewDecoder(r.Body).Decode(&updatedCustomer)
// 	customers := models.UpdateCustomer(cusId, updatedCustomer)
// 	res, e := json.Marshal(customers)
// 	if e != nil {
// 		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
// 		http.Error(w, string(resp), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }

// func DeleteCustomer(w http.ResponseWriter, r *http.Request) {

// 	queryParams := mux.Vars(r)
// 	customers := models.DeleteCustomer(queryParams["cusId"])
// 	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
// 	res, e := json.Marshal(customers)
// 	if e != nil {
// 		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
// 		http.Error(w, string(resp), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }
var db = config.GetDB()

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers []models.Customers
	e := db.Find(&customers).Error
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
		return
	}
	err := json.NewEncoder(w).Encode(customers)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	var customers models.Customers
	// db.Raw("select id, date_of_order, order_details_id, customer_id, supplier_id, status from Orders where id=?",1).Scan(&order)
	// data,_:= json.Marshal(&order)
	// fmt.Fprint(w, string(data))
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["cusId"])
	db.First(&customers, queryParams["cusId"])
	e := json.NewEncoder(w).Encode(customers)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
	}
}

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers models.Customers
	json.NewDecoder(r.Body).Decode(&customers)
	db.Create(&customers)
	json.NewEncoder(w).Encode(customers)
	// fmt.Println("Hello")
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var customers models.Customers
	var updatedCustomer models.Customers
	queryParams := mux.Vars(r)
	cusId := queryParams["cusId"]
	db.First(&customers, cusId)
	json.NewDecoder(r.Body).Decode(&updatedCustomer)
	db.Model(&customers).Where("customer_id=?", cusId).Updates(&updatedCustomer)
	json.NewEncoder(w).Encode(&customers)

}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	// var customer models.Customers
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["cusId"])
	er := db.Exec("delete from customers where customer_id=?", queryParams["cusId"]).Error
	if er != nil {
		json.NewEncoder(w).Encode("{Status:201, Message: Internal Error}")
	}
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

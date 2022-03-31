package main

import (

	//     "log"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customers struct {
	// gorm.Model
	CUSTOMER_ID uint   `json:"customer_id"`
	FIRST_NAME  string `json:"first_name"`
	LAST_NAME   string `json:"last_name"`
	ADDRESS     string `json:"address"`
	PHONE       uint   `json:"phone"`
	EMAIL       string `json:"email"`
	STAFF_ID    uint   `json:"staff_id"`
}

/*
var db *gorm.DB
var err error

func InitDB() {
	db, err = gorm.Open(sqlite.Open("./INVENTORY_MANAGEMENT.db"), &gorm.Config{})
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
*/

<<<<<<< HEAD
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
=======
func getAllCustomer(w http.ResponseWriter, r *http.Request) {
>>>>>>> 18a19f1 (added supplier API)
	w.Header().Set("Content-Type", "application/json")
	var customers []Customers
	e := db.Find(&customers).Error
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	e = json.NewEncoder(w).Encode(customers)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func getCustomerByID(w http.ResponseWriter, r *http.Request) {
	var customers Customers
	// db.Raw("select id, date_of_order, order_details_id, customer_id, supplier_id, status from Orders where id=?",1).Scan(&order)
	// data,_:= json.Marshal(&order)
	// fmt.Fprint(w, string(data))
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["cusId"])
	db.First(&customers, queryParams["cusId"])
	e := json.NewEncoder(w).Encode(customers)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers Customers
	json.NewDecoder(r.Body).Decode(&customers)
	db.Create(&customers)
	json.NewEncoder(w).Encode(customers)
	// fmt.Println("Hello")
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	var customers Customers
	var updatedCustomer Customers
	queryParams := mux.Vars(r)
	cusId := queryParams["cusId"]
	db.First(&customers, cusId)
	json.NewDecoder(r.Body).Decode(&updatedCustomer)
	db.Model(&customers).Where("customer_id=?", cusId).Updates(&updatedCustomer)
	json.NewEncoder(w).Encode(&customers)

}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customers
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["cusId"])
	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
	db.Delete(&customer, queryParams["cusId"])
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

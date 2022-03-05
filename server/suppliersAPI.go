package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Suppliers struct {
	// gorm.Model
	SUPPLIER_ID uint   `json:"supplier_id"`
	NAME        string `json:"name"`
	ADDRESS     string `json:"address"`
	PHONE       uint   `json:"phone"`
	EMAIL       string `json:"email"`
	//OTHER_DETAILS string `json:"staff_id"`
	// In dataBase, we have OTHER DETAILS....We need to sort this column name
	//so that we can put and post this other_details as well :)
}

func getAllSuppliers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var suppliers []Suppliers
	e := db.Find(&suppliers).Error
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	e = json.NewEncoder(w).Encode(suppliers)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func getSupplierByID(w http.ResponseWriter, r *http.Request) {
	var suppliers Suppliers
	// db.Raw("select id, date_of_order, order_details_id, customer_id, supplier_id, status from Orders where id=?",1).Scan(&order)
	// data,_:= json.Marshal(&order)
	// fmt.Fprint(w, string(data))
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["supId"])
	db.First(&suppliers, queryParams["supId"])
	e := json.NewEncoder(w).Encode(suppliers)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func addSupplier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var supplier Suppliers
	json.NewDecoder(r.Body).Decode(&supplier)
	db.Create(&supplier)
	json.NewEncoder(w).Encode(supplier)
	// fmt.Println("Hello")
}

func updateSupplier(w http.ResponseWriter, r *http.Request) {
	var supplier Suppliers
	var updatedSupplier Suppliers
	queryParams := mux.Vars(r)
	cusId := queryParams["supId"]
	db.First(&supplier, cusId)
	json.NewDecoder(r.Body).Decode(&updatedSupplier)
	db.Model(&supplier).Where("customer_id=?", cusId).Updates(&updatedSupplier)
	json.NewEncoder(w).Encode(&supplier)

}

func deleteSupplier(w http.ResponseWriter, r *http.Request) {
	var supplier Suppliers
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["supId"])
	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
	db.Delete(&supplier, queryParams["supId"])
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

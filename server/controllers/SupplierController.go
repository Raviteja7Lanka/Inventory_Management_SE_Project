package controllers

import (
	"apis/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// var db *gorm.DB

// func GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
// 	// w.Header().Set("Content-Type", "application/json")
// 	suppliers := models.GetAllSuppliers()
// 	res, e := json.Marshal(suppliers)
// 	if e != nil {
// 		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
// 		http.Error(w, string(resp), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func GetSupplierByID(w http.ResponseWriter, r *http.Request) {

// 	queryParams := mux.Vars(r)

// 	supplier := models.GetSupplierByID(queryParams["supId"])
// 	res, e := json.Marshal(supplier)
// 	if e != nil {
// 		http.Error(w, string(e.Error()), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func AddSupplier(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	supplier := &models.Suppliers{}
// 	json.NewDecoder(r.Body).Decode(&supplier)
// 	newCust := models.AddSupplier(supplier)
// 	res, e := json.Marshal(newCust)
// 	if e != nil {
// 		http.Error(w, string(e.Error()), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// 	// fmt.Println("Hello")
// }

// func UpdateSupplier(w http.ResponseWriter, r *http.Request) {

// 	updatedSupplier := &models.Suppliers{}
// 	queryParams := mux.Vars(r)
// 	supId := queryParams["supId"]
// 	json.NewDecoder(r.Body).Decode(&updatedSupplier)
// 	staff := models.UpdateSupplier(supId, updatedSupplier)
// 	res, e := json.Marshal(staff)
// 	if e != nil {
// 		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
// 		http.Error(w, string(resp), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }

// func DeleteSupplier(w http.ResponseWriter, r *http.Request) {

// 	queryParams := mux.Vars(r)
// 	supplier := models.DeleteSupplier(queryParams["supId"])
// 	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
// 	res, e := json.Marshal(supplier)
// 	if e != nil {
// 		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
// 		http.Error(w, string(resp), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }

func GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var suppliers []models.Suppliers
	e := db.Find(&suppliers).Error
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
		return
	}
	e = json.NewEncoder(w).Encode(suppliers)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
	}
}

func GetSupplierByID(w http.ResponseWriter, r *http.Request) {
	var suppliers models.Suppliers
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["supId"])
	db.First(&suppliers, queryParams["supId"])
	e := json.NewEncoder(w).Encode(suppliers)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
	}
}

func AddSupplier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var supplier models.Suppliers
	json.NewDecoder(r.Body).Decode(&supplier)
	db.Create(&supplier)
	json.NewEncoder(w).Encode(supplier)
}

func UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	var supplier models.Suppliers
	var updatedSupplier models.Suppliers
	queryParams := mux.Vars(r)
	cusId := queryParams["supId"]
	db.First(&supplier, cusId)
	json.NewDecoder(r.Body).Decode(&updatedSupplier)
	db.Model(&supplier).Where("customer_id=?", cusId).Updates(&updatedSupplier)
	json.NewEncoder(w).Encode(&supplier)

}

func DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	// var supplier models.Suppliers
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["supId"])
	er := db.Exec("delete from suppliers where supplier_id=?", queryParams["supId"]).Error
	if er != nil {
		json.NewEncoder(w).Encode("{Status:201, Message: Internal Error}")
	}
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

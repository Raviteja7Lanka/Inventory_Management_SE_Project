package controllers

import (
	"apis/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// var db *gorm.DB

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	customers := models.GetAllCustomers()
	res, e := json.Marshal(customers)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	// db.Raw("select id, date_of_order, order_details_id, customer_id, supplier_id, status from Orders where id=?",1).Scan(&order)
	// data,_:= json.Marshal(&order)
	// fmt.Fprint(w, string(data))
	queryParams := mux.Vars(r)
	// fmt.Println(queryParams["cusId"])
	// db.First(&customers, queryParams["cusId"])
	customers := models.GetCustomerByID(queryParams["cusId"])
	res, e := json.Marshal(customers)
	if e != nil {
		http.Error(w, string(e.Error()), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	customers := &models.Customers{}
	json.NewDecoder(r.Body).Decode(&customers)
	newCust := models.AddCustomer(customers)
	res, e := json.Marshal(newCust)
	if e != nil {
		http.Error(w, string(e.Error()), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// fmt.Println("Hello")
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {

	updatedCustomer := &models.Customers{}
	queryParams := mux.Vars(r)
	cusId := queryParams["cusId"]
	json.NewDecoder(r.Body).Decode(&updatedCustomer)
	customers := models.UpdateCustomer(cusId, updatedCustomer)
	res, e := json.Marshal(customers)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {

	queryParams := mux.Vars(r)
	customers := models.DeleteCustomer(queryParams["cusId"])
	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
	res, e := json.Marshal(customers)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

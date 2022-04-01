package controllers

import (
	"apis/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//var db *gorm.DB

func GetAllPayments(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	Payments := models.GetAllPayments()
	res, e := json.Marshal(Payments)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPaymentByID(w http.ResponseWriter, r *http.Request) {
	// db.Raw("select id, date_of_order, order_details_id, Order_id, supplier_id, status from Orders where id=?",1).Scan(&order)
	// data,_:= json.Marshal(&order)
	// fmt.Fprint(w, string(data))
	queryParams := mux.Vars(r)
	// fmt.Println(queryParams["ordId"])
	// db.First(&Orders, queryParams["ordId"])
	Orders := models.GetPaymentByID(queryParams["payId"])
	res, e := json.Marshal(Orders)
	if e != nil {
		http.Error(w, string(e.Error()), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AddPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Payments := &models.Payments{}
	json.NewDecoder(r.Body).Decode(&Payments)
	newCust := models.AddPayment(Payments)
	res, e := json.Marshal(newCust)
	if e != nil {
		http.Error(w, string(e.Error()), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// fmt.Println("Hello")
}

func UpdatePayment(w http.ResponseWriter, r *http.Request) {

	updatedPayment := &models.Payments{}
	queryParams := mux.Vars(r)
	payId := queryParams["payId"]
	json.NewDecoder(r.Body).Decode(&updatedPayment)
	payments := models.UpdatePayment(payId, updatedPayment)
	res, e := json.Marshal(payments)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeletePayment(w http.ResponseWriter, r *http.Request) {

	queryParams := mux.Vars(r)
	Payments := models.DeletePayment(queryParams["payId"])
	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
	res, e := json.Marshal(Payments)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

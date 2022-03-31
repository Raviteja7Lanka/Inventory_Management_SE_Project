package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Payments struct {
	// gorm.Model
	BILL_NUMBER   uint   `json:"bill_number"`
	PAYMENT_TYPE  string `json:"payment_type"`
	OTHER_DETAILS string `json:"other_details"`
}

func getAllPayments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payments []Payments
	e := db.Find(&payments).Error
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	e = json.NewEncoder(w).Encode(payments)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func getPaymentsByID(w http.ResponseWriter, r *http.Request) {
	var payments Payments
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["payId"])
	db.First(&payments, queryParams["payId"])
	e := json.NewEncoder(w).Encode(payments)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func addPayments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payments Payments
	json.NewDecoder(r.Body).Decode(&payments)
	db.Create(&payments)
	json.NewEncoder(w).Encode(payments)
}

func updatePayments(w http.ResponseWriter, r *http.Request) {
	var payments Payments
	var updatedPayments Payments
	queryParams := mux.Vars(r)
	payId := queryParams["payId"]
	db.First(&payments, payId)
	json.NewDecoder(r.Body).Decode(&updatedPayments)
	db.Model(&payments).Where("customer_id=?", payId).Updates(&updatedPayments)
	json.NewEncoder(w).Encode(&payments)

}

func deletePayments(w http.ResponseWriter, r *http.Request) {
	var payments Payments
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["payId"])
	db.Delete(&payments, queryParams["payId"])
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

package main

import (
	"fmt"
	//     "log"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Order_details struct {
	// gorm.Model
	OrderDetailsId uint    `json:"order_details_id"`
	UnitPrice      string  `json:"UnitPrice"`
	Size           uint    `json:"size"` //blob is there in schema
	Quantity       uint    `json:"quantity"`
	Discount       float64 `json:"discount"`
	Total          float64 `json:"total"`
	Date           string  `json:"date"`
	ProductId      string  `json:"product_id"`
	OrderId        string  `json:"order_id"`
	BillNumber     uint    `json:"billnumber"`
}

func getAllOrderDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order_details []Order_details
	e := db.Find(&order_details).Error
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	e = json.NewEncoder(w).Encode(order_details)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}

}

func getOrderDetailsByID(w http.ResponseWriter, r *http.Request) {
	var order_details Order_details
	// db.Raw("select id, date_of_order, order_details_id, customer_id, supplier_id, status from Orders where id=?",1).Scan(&order)
	// data,_:= json.Marshal(&order)
	// fmt.Fprint(w, string(data))
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["order_details_id"])
	db.First(&order_details, queryParams["order_details_id"])
	e := json.NewEncoder(w).Encode(order_details)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}
func addOrderDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order_details Order_details
	json.NewDecoder(r.Body).Decode(&order_details)
	db.Create(&order_details)
	json.NewEncoder(w).Encode(order_details)
	// fmt.Println("Hello")
}
func updateOrderDetails(w http.ResponseWriter, r *http.Request) {
	var order_details Order_details
	var updatedOrderDetails Order_details
	queryParams := mux.Vars(r)
	ordId := queryParams["order_details_id"]
	db.First(&order_details, ordId)
	json.NewDecoder(r.Body).Decode(&updatedOrderDetails)
	db.Model(&order_details).Where("order_details_id=?", ordId).Updates(&updatedOrderDetails)
	json.NewEncoder(w).Encode(&order_details)

}

func deleteOrderDetails(w http.ResponseWriter, r *http.Request) {
	var order_details Order_details
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["order_details_id"])
	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
	db.Delete(&order_details, queryParams["order_details_id"])
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

// func handleRequests() {
//     http.HandleFunc("/customer/orders/all", getAllCustomerOrders)
//     http.HandleFunc("/customer/order/{ordId}", getCustomerOrderByID)
//     http.HandleFunc("/customer/orders/add/", addCustomerOrder)
//     http.HandleFunc("/customer/orders/update/", updateCustomerOrder)
//     log.Fatal(http.ListenAndServe(":8085", nil))
//     fmt.Println("Hello world")
// }

// func main() {
//     fmt.Println("Hello world")
//     //handleRequests()
// }

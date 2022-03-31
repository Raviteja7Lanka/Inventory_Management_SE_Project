// package controllers

// import (
// 	"/usr/local/go/src/apis/server/models"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func getAllCustomers(w http.ResponseWriter, r *http.Request) {
// 	// w.Header().Set("Content-Type", "application/json")
// 	customers := models.getAllCustomers(w, r)
// 	res, e := json.Marshal(customers)
// 	if e != nil {
// 		resp, _ := json.Marshal(map[string]string{"error": e.message})
// 		http.Error(w, string(resp), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func getCustomerByID(w http.ResponseWriter, r *http.Request) {
// 	var customers Customers
// 	// db.Raw("select id, date_of_order, order_details_id, customer_id, supplier_id, status from Orders where id=?",1).Scan(&order)
// 	// data,_:= json.Marshal(&order)
// 	// fmt.Fprint(w, string(data))
// 	// queryParams := mux.Vars(r)
// 	// fmt.Println(queryParams["cusId"])
// 	// db.First(&customers, queryParams["cusId"])
// 	customers = models.getCustomerByID(&customers, queryParams["cusId"])
// 	res, e := json.Marshal(customers)
// 	if e != nil {
// 		resp, _ := json.Marshal(map[string]string{"error": e.message})
// 		http.Error(w, string(resp), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func addCustomer(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var customers Customers
// 	json.NewDecoder(r.Body).Decode(&customers)
// 	db.Create(&customers)
// 	json.NewEncoder(w).Encode(customers)
// 	// fmt.Println("Hello")
// }

// func updateCustomer(w http.ResponseWriter, r *http.Request) {
// 	var customers Customers
// 	var updatedCustomer Customers
// 	queryParams := mux.Vars(r)
// 	cusId := queryParams["cusId"]
// 	db.First(&customers, cusId)
// 	json.NewDecoder(r.Body).Decode(&updatedCustomer)
// 	db.Model(&customers).Where("customer_id=?", cusId).Updates(&updatedCustomer)
// 	json.NewEncoder(w).Encode(&customers)

// }

// func deleteCustomer(w http.ResponseWriter, r *http.Request) {
// 	var customer Customers
// 	queryParams := mux.Vars(r)
// 	fmt.Println(queryParams["cusId"])
// 	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
// 	db.Delete(&customer, queryParams["cusId"])
// 	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
// }

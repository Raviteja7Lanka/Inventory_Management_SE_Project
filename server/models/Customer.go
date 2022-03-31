package models

import (
	// "fmt"
	//     "log"
	// "encoding/json"
	// "net/http"
	// "apis/server/config"
	// "github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Customers struct {
	gorm.Model
	CUSTOMER_ID uint   `json:"customer_id"`
	FIRST_NAME  string `json:"first_name"`
	LAST_NAME   string `json:"last_name"`
	ADDRESS     string `json:"address"`
	PHONE       uint   `json:"phone"`
	EMAIL       string `json:"email"`
	STAFF_ID    uint   `json:"staff_id"`
}

// var db *gorm.DB
// var err error

// // func sendErr(w http.ResponseWriter, code int, message string) {
// // 	resp, _ := json.Marshal(map[string]string{"error": message})
// // 	http.Error(w, string(resp), code)
// // }
// func init() {
// 	config.ConnectDB()
// 	db = config.GetDB()
// 	userdb.AutoMigrate(&Customers{})
// }
// func getAllCustomers() []Customers {

// 	var customers []Customers
// 	db.Find(&customers)
// 	// if e != nil {
// 	// 	sendErr(w, http.StatusInternalServerError, e.Error())
// 	// 	return customers
// 	// }
// 	// if e != nil {
// 	// 	sendErr(w, http.StatusInternalServerError, e.Error())
// 	// }
// 	return customers
// }

// func getCustomerByID(string cusId) Customers {
// 	var customers Customers
// 	db.First(&customers,cusId)
// 	e := json.NewEncoder(w).Encode(customers)
// 	if e != nil {
// 		sendErr(w, http.StatusInternalServerError, err.Error())
// 	}
// 	return customers
// }

// func addCustomer(w http.ResponseWriter, r *http.Request) Customers {
// 	w.Header().Set("Content-Type", "application/json")
// 	var customers Customers
// 	json.NewDecoder(r.Body).Decode(&customers)
// 	db.Create(&customers)
// 	json.NewEncoder(w).Encode(customers)
// 	// fmt.Println("Hello")
// 	return customers
// }

// func updateCustomer(w http.ResponseWriter, r *http.Request) Customers {
// 	var customers Customers
// 	var updatedCustomer Customers
// 	queryParams := mux.Vars(r)
// 	cusId := queryParams["cusId"]
// 	db.First(&customers, cusId)
// 	json.NewDecoder(r.Body).Decode(&updatedCustomer)
// 	db.Model(&customers).Where("customer_id=?", cusId).Updates(&updatedCustomer)
// 	json.NewEncoder(w).Encode(&customers)
// 	return customers
// }

// func deleteCustomer(w http.ResponseWriter, r *http.Request) Customers {
// 	var customer Customers
// 	queryParams := mux.Vars(r)
// 	fmt.Println(queryParams["cusId"])
// 	db.Delete(&customer, queryParams["cusId"])
// 	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
// 	return customer
// }

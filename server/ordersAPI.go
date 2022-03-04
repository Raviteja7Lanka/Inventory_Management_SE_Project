package main

import (
	"fmt"
	//     "log"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Orders struct {
	// gorm.Model
	OrderId        uint   `json:"order_id"`
	DateOfOrder    string `json:"date_of_order"`
	OrderDetailsId string `json:"order_details_id"`
	CustomerId     uint   `json:"customer_id"`
	SupplierId     uint   `json:"supplier_id"`
	Status         string `json:"status"`
}

//
/*
"CUSTOMER_ID"	INTEGER,
	"FIRST_NAME"	TEXT,
	"LAST_NAME"	REAL,
	"ADDRESS"	TEXT,
	"PHONE"	INTEGER,
	"EMAIL"	TEXT,
	"STAFF_ID"	INTEGER,

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
*/

//

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

func getAllCustomerOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var orders []Orders
	e := db.Find(&orders).Error
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	e = json.NewEncoder(w).Encode(orders)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
	// data,_:= json.Marshal(&orders)
	// fmt.Fprint(w, string(data))

	// fmt.Println("Hello world")
	// log.Fatal("Hello")
	// db.Raw("select id, date_of_order, order_details_id, customer_id, supplier_id, status from Orders where id=?",1).Scan(&orders)
	// data,_:= json.Marshal(&orders)
	// fmt.Fprint(w, string(data))

}

/*
// CUSTOMER Table

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
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

//end
*/

// SUPPLIER Table

//end

func getCustomerOrderByID(w http.ResponseWriter, r *http.Request) {
	var order Orders
	// db.Raw("select id, date_of_order, order_details_id, customer_id, supplier_id, status from Orders where id=?",1).Scan(&order)
	// data,_:= json.Marshal(&order)
	// fmt.Fprint(w, string(data))
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["ordId"])
	db.First(&order, queryParams["ordId"])
	e := json.NewEncoder(w).Encode(order)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

/*
//
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
*/

//

func addCustomerOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Orders
	json.NewDecoder(r.Body).Decode(&order)
	db.Create(&order)
	json.NewEncoder(w).Encode(order)
	// fmt.Println("Hello")
}

/*
//
func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers Customers
	json.NewDecoder(r.Body).Decode(&customers)
	db.Create(&customers)
	json.NewEncoder(w).Encode(customers)
	// fmt.Println("Hello")
}
*/

//
func updateCustomerOrder(w http.ResponseWriter, r *http.Request) {
	var order Orders
	var updatedOrder Orders
	queryParams := mux.Vars(r)
	ordId := queryParams["ordId"]
	db.First(&order, ordId)
	json.NewDecoder(r.Body).Decode(&updatedOrder)
	db.Model(&order).Where("order_id=?", ordId).Updates(&updatedOrder)
	json.NewEncoder(w).Encode(&order)

}

/*
//
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
*/

//

func deleteCustomerOrder(w http.ResponseWriter, r *http.Request) {
	var order Orders
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["ordId"])
	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
	db.Delete(&order, queryParams["ordId"])
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

/*
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customers
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["cusId"])
	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
	db.Delete(&customer, queryParams["cusId"])
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}
*/

//

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

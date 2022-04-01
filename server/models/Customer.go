package models

import (
	"apis/config"
	"fmt"

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

var db *gorm.DB
var err error

// func sendErr(w http.ResponseWriter, code int, message string) {
// 	resp, _ := json.Marshal(map[string]string{"error": message})
// 	http.Error(w, string(resp), code)
// }
func init() {
	config.ConnectSqlite()
	db = config.GetDB()
	db.AutoMigrate(&Customers{})
}
func GetAllCustomers() []Customers {

	var customers []Customers
	e := db.Find(&customers)
	// if e != nil {
	// 	sendErr(w, http.StatusInternalServerError, e.Error())
	// 	return customers
	// }
	if e != nil {
		fmt.Println("Erroring getting all Customers")
	}
	return customers
}

func GetCustomerByID(cusId string) Customers {
	var customers Customers
	e := db.First(&customers, cusId)
	if e != nil {
		fmt.Println("Error finding the customer")
	}
	return customers
}

func AddCustomer(customers *Customers) *Customers {

	e := db.Create(&customers)
	// fmt.Println("Hello")
	if e != nil {
		fmt.Println("Error creating new customer")
	}
	return customers
}

func UpdateCustomer(cusId string, updatedCustomer *Customers) Customers {
	var customers Customers
	db.First(&customers, cusId)
	db.Model(&customers).Where("customer_id=?", cusId).Updates(&updatedCustomer)
	return customers
}

func DeleteCustomer(cusId string) Customers {
	var customer Customers
	db.Delete(&customer, cusId)
	// json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
	return customer
}

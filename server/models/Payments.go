package models

import (
	"apis/config"

	"gorm.io/gorm"
)

type Payments struct {
	gorm.Model
	BILL_NUMBER   string `json:"bill_number"`
	PAYMENT_TYPE  string `json:"payment_type"`
	OTHER_DETAILS string `json:"other_details"`
}

func init() {
	config.ConnectSqlite()
	db := config.GetDB()
	db.AutoMigrate(&Payments{})
}

// func GetAllPayments() []Payments {

// 	var payments []Payments
// 	e := db.Find(&payments)
// 	// if e != nil {
// 	// 	sendErr(w, http.StatusInternalServerError, e.Error())
// 	// 	return Orders
// 	// }
// 	if e != nil {
// 		fmt.Println("Erroring getting all Payments")
// 	}
// 	return payments
// }

// func GetPaymentByID(cusId string) Payments {
// 	var payments Payments
// 	e := db.First(&payments, cusId)
// 	if e != nil {
// 		fmt.Println("Error finding the Order")
// 	}
// 	return payments
// }

// func AddPayment(payments *Payments) *Payments {

// 	e := db.Create(&payments)
// 	// fmt.Println("Hello")
// 	if e != nil {
// 		fmt.Println("Error creating new Payment")
// 	}
// 	return payments
// }

// func UpdatePayment(cusId string, updatedPayment *Payments) Payments {
// 	var payments Payments
// 	db.First(&payments, cusId)
// 	db.Model(&payments).Where("Order_id=?", cusId).Updates(&updatedPayment)
// 	return payments
// }

// func DeletePayment(cusId string) Payments {
// 	var payments Payments
// 	db.Delete(&payments, cusId)
// 	// json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
// 	return payments
// }

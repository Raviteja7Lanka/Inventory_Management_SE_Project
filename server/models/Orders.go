package models

import (
	"apis/config"
	"fmt"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	OrderId        uint   `json:"order_id"`
	DateOfOrder    string `json:"date_of_order"`
	OrderDetailsId string `json:"order_details_id"`
	OrderrId       uint   `json:"Order_id"`
	SupplierId     uint   `json:"supplier_id"`
	Status         string `json:"status"`
}

func init() {
	config.ConnectSqlite()
	db = config.GetDB()
	db.AutoMigrate(&Orders{})
}
func GetAllOrders() []Orders {

	var orders []Orders
	e := db.Find(&orders)
	// if e != nil {
	// 	sendErr(w, http.StatusInternalServerError, e.Error())
	// 	return Orders
	// }
	if e != nil {
		fmt.Println("Erroring getting all Orders")
	}
	return orders
}

func GetOrderByID(cusId string) Orders {
	var orders Orders
	e := db.First(&orders, cusId)
	if e != nil {
		fmt.Println("Error finding the Order")
	}
	return orders
}

func AddOrder(orders *Orders) *Orders {

	e := db.Create(&orders)
	// fmt.Println("Hello")
	if e != nil {
		fmt.Println("Error creating new Order")
	}
	return orders
}

func UpdateOrder(cusId string, updatedOrder *Orders) Orders {
	var orders Orders
	db.First(&orders, cusId)
	db.Model(&orders).Where("Order_id=?", cusId).Updates(&updatedOrder)
	return orders
}

func DeleteOrder(cusId string) Orders {
	var order Orders
	db.Delete(&order, cusId)
	// json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
	return order
}

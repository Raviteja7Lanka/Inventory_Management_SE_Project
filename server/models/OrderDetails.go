package models

import (
	"apis/config"
)

type OrderDetails struct {
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

func init() {
	config.ConnectSqlite()
	db := config.GetDB()
	db.AutoMigrate(&OrderDetails{})
}

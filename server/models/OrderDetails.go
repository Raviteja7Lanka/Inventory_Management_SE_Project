package models

import (
	"apis/config"

	"gorm.io/gorm"
)

type OrderDetails struct {
	gorm.Model
	OrderDetailsId string `json:"order_details_id"`
	UnitPrice      string `json:"unit_price"`
	Size           uint   `json:"size"` //blob is there in schema
	Quantity       uint   `json:"quantity"`
	Discount       uint   `json:"discount"`
	Total          uint   `json:"total"`
	DateOfOrder    string `json:"date_of_order"`
	ProductId      string `json:"product_id"`
	OrderId        string `json:"order_id"`
	BillNumber     string `json:"billnumber"`
}

func init() {
	config.ConnectSqlite()
	db := config.GetDB()
	db.AutoMigrate(&OrderDetails{})
}

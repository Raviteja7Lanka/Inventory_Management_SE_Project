package models

import (
	"apis/config"
	"fmt"

	"gorm.io/gorm"
)

type Suppliers struct {
	gorm.Model
	SUPPLIER_ID uint   `json:"supplier_id"`
	NAME        string `json:"name"`
	ADDRESS     string `json:"address"`
	PHONE       uint   `json:"phone"`
	EMAIL       string `json:"email"`
}

func init() {
	config.ConnectSqlite()
	db = config.GetDB()
	db.AutoMigrate(&Orders{})
}

func GetAllSuppliers() []Suppliers {
	var suppliers []Suppliers
	e := db.Find(&suppliers)

	if e != nil {
		fmt.Println("Error in  getting all Suppliers")
	}
	return suppliers
}

func GetSupplierByID(supId string) Suppliers {
	var suppliers Suppliers
	e := db.First(&suppliers, supId)
	if e != nil {
		fmt.Println("Error finding the requested Supplier")
	}
	return suppliers
}

func AddSupplier(suppliers *Suppliers) *Suppliers {

	e := db.Create(&suppliers)
	// fmt.Println("Hello")
	if e != nil {
		fmt.Println("Error creating a new supplier")
	}
	return suppliers
}

func UpdateSupplier(supId string, updateSupplier *Suppliers) Suppliers {
	var suppliers Suppliers
	db.First(&suppliers, supId)
	db.Model(&suppliers).Where("Product_id=?", supId).Updates(&updateSupplier)
	return suppliers
}

func DeleteSupplier(supId string) Suppliers {
	var suppliers Suppliers
	db.Delete(&suppliers, supId)
	// json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
	return suppliers
}

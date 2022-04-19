package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Warehouses struct {
	gorm.Model
	WAREHOUSE_ID string `json:"warehouse_id"`
	NAME         string `json:"name"`
	LOCATION     string `json:"location"`
	DES          string `json:"des"`
	STATUS       string `json:"status"`
}

// func sendErr(w http.ResponseWriter, code int, message string) {
// 	resp, _ := json.Marshal(map[string]string{"error": message})
// 	http.Error(w, string(resp), code)
// }

func GetAllWarehouses() []Warehouses {

	var warehouses []Warehouses
	e := db.Find(&warehouses)
	// if e != nil {
	// 	sendErr(w, http.StatusInternalServerError, e.Error())
	// 	return customers
	// }
	if e != nil {
		fmt.Println("Erroring getting all Warehouses")
	}
	return warehouses
}

func GetWarehouseByID(wareId string) Warehouses {
	var warehouses Warehouses
	e := db.First(&warehouses, wareId)
	if e != nil {
		fmt.Println("Error finding the warehouses")
	}
	return warehouses
}

func AddWarehouse(warehouses *Warehouses) *Warehouses {

	e := db.Create(&warehouses)
	// fmt.Println("Hello")
	if e != nil {
		fmt.Println("Error creating new warehouses")
	}
	return warehouses
}

func UpdateWarehouse(wareId string, updatedWarehouse *Warehouses) Warehouses {
	var warehouses Warehouses
	db.First(&warehouses, wareId)
	db.Model(&warehouses).Where("warehouse_id=?", wareId).Updates(&updatedWarehouse)
	return warehouses
}

func DeleteWarehouse(wareId string) Warehouses {
	var warehouses Warehouses
	db.Delete(&warehouses, wareId)
	// json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
	return warehouses
}

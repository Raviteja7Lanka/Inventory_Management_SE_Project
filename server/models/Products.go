package models

import (
	"apis/config"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	PRODUCTS_ID           string `json:"products_id"`
	PRODUCT_NAME          string `json:"product_name"`
	PRODUCT_DESCR_IP_TION string `json:"product_descr_ip_tion"`
	PRODUCT_UNIT          string `json:"product_unit"`
	PRODUCT_PRICE         string `json:"product_price"`
	PRODUCT_QUANTITY      string `json:"product_quantity"`
	PRODUCT_STATUS        string `json:"product_status"`
	OTHER_DETAILS         string `json:"other_details"`
	SUPPLIER_ID           string `json:"supplier_id"`
	CATEGORY_ID           string `json:"category_id"`
	//
}

func init() {
	config.ConnectSqlite()
	db := config.GetDB()
	db.AutoMigrate(&Products{})
}

// func GetAllProducts() []Products {
// 	var products []Products
// 	e := db.Find(&products)
// 	// if e != nil {
// 	// 	sendErr(w, http.StatusInternalServerError, err.Error())
// 	// 	return
// 	// }
// 	//e = json.NewEncoder(w).Encode(products)
// 	if e != nil {
// 		fmt.Println("Error in  getting all Products")
// 	}
// 	return products
// }

// func GetProductByID(prodId string) Products {
// 	var products Products
// 	e := db.First(&products, prodId)
// 	if e != nil {
// 		fmt.Println("Error finding the Product")
// 	}
// 	return products
// }

// func AddProduct(products *Products) *Products {

// 	e := db.Create(&products)
// 	// fmt.Println("Hello")
// 	if e != nil {
// 		fmt.Println("Error creating a new Product")
// 	}
// 	return products
// }

// func UpdateProduct(prodId string, updateProduct *Products) Products {
// 	var products Products
// 	db.First(&products, prodId)
// 	db.Model(&products).Where("Product_id=?", prodId).Updates(&updateProduct)
// 	return products
// }

// func DeleteProduct(prodId string) Products {
// 	var products Products
// 	db.Delete(&products, prodId)
// 	// json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
// 	return products
// }

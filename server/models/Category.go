package models

import (
	"apis/config"
	"fmt"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CATEGORY_ID   uint   `json:"category_id"`
	CATEGORY_NAME string `json:"category_name"`
	DESCRIPTION   string `json:"description"`
}

// func sendErr(w http.ResponseWriter, code int, message string) {
// 	resp, _ := json.Marshal(map[string]string{"error": message})
// 	http.Error(w, string(resp), code)
// }
func init() {
	config.ConnectSqlite()
	db = config.GetDB()
	db.AutoMigrate(&Category{})
}
func GetAllCategories() []Category {

	var category []Category
	e := db.Find(&category)
	// if e != nil {
	// 	sendErr(w, http.StatusInternalServerError, e.Error())
	// 	return customers
	// }
	if e != nil {
		fmt.Println("Erroring getting all Categories")
	}
	return category
}

func GetCategoryByID(catId string) Category {
	var category Category
	e := db.First(&category, catId)
	if e != nil {
		fmt.Println("Error finding the category")
	}
	return category
}

func AddCategories(category *Category) *Category {

	e := db.Create(&category)
	// fmt.Println("Hello")
	if e != nil {
		fmt.Println("Error creating new category")
	}
	return category
}

func UpdateCategories(catId string, updatedCategory *Category) Category {
	var category Category
	db.First(&category, catId)
	db.Model(&category).Where("category_id=?", catId).Updates(&updatedCategory)
	return category
}

func DeleteCategories(catId string) Category {
	var category Category
	db.Delete(&category, catId)
	// json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
	return category
}

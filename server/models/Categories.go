package models

import (
	"apis/config"

	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	CATEGORY_ID  string `json:"category_id"`
	WAREHOUSE_ID string `json:"category_id"`
	NAME         string `json:"name"`
	DESCRIPTION  string `json:"description"`
}

// func sendErr(w http.ResponseWriter, code int, message string) {
// 	resp, _ := json.Marshal(map[string]string{"error": message})
// 	http.Error(w, string(resp), code)
// }
func init() {
	config.ConnectSqlite()
	db := config.GetDB()
	db.AutoMigrate(&Categories{})
}

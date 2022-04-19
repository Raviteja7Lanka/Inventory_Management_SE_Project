package models

import (
	"apis/config"

	"gorm.io/gorm"
)

type Warehouses struct {
	gorm.Model
	WAREHOUSE_ID string `json:"warehouse_id"`
	NAME         string `json:"name"`
	LOCATION     string `json:"location"`
	DESCRIPTION  string `json:"description"`
	STATUS       string `json:"status"`
	CAPACITY     uint   `json:"capacity"`
}

// func sendErr(w http.ResponseWriter, code int, message string) {
// 	resp, _ := json.Marshal(map[string]string{"error": message})
// 	http.Error(w, string(resp), code)
// }
func init() {
	config.ConnectSqlite()
	db := config.GetDB()
	db.AutoMigrate(&Warehouses{})
}

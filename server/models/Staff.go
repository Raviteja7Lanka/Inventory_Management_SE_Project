package models

import (
	"apis/config"

	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	STAFFS_ID  string `json:"staffs_id"`
	FIRST_NAME string `json:"first_name"`
	LAST_NAME  string `json:"last_name"`
	ADDRESS    string `json:"address"`
	PHONE      string `json:"phone"`
	EMAIL      string `json:"email"`
	USERNAME   string `json:"username"`
	PASSWORD   string `json:"password"`
}

func init() {
	config.ConnectSqlite()
	db := config.GetDB()
	db.AutoMigrate(&Staff{})
}

// func GetAllStaff() []Staff {
// 	var staff []Staff
// 	e := db.Find(&staff)

// 	if e != nil {
// 		fmt.Println("Error in  getting all Staff")
// 	}
// 	return staff
// }

// func GetStaffByID(staffId string) Staff {
// 	var staff Staff
// 	e := db.First(&staff, staffId)
// 	if e != nil {
// 		fmt.Println("Error finding the requested Staff")
// 	}
// 	return staff
// }

// func AddStaff(staff *Staff) *Staff {

// 	e := db.Create(&staff)
// 	// fmt.Println("Hello")
// 	if e != nil {
// 		fmt.Println("Error creating a new Staff")
// 	}
// 	return staff
// }

// func UpdateStaff(staffId string, UpdateStaff *Staff) Staff {
// 	var staff Staff
// 	db.First(&staff, staffId)
// 	db.Model(&staff).Where("Product_id=?", staffId).Updates(&UpdateStaff)
// 	return staff
// }

// func DeleteStaff(staffId string) Staff {
// 	var staff Staff
// 	db.Delete(&staff, staffId)
// 	// json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
// 	return staff
// }

package models

import (
	"apis/config"
	"fmt"

	"gorm.io/gorm"
)

type Roles struct {
	gorm.Model
	ROLE_ID     uint   `json:"role_id"`
	ROLE_NAME   string `json:"role_name"`
	DESCRIPTION string `json:"description"`
	//

}

func init() {
	config.ConnectSqlite()
	db = config.GetDB()
	db.AutoMigrate(&Orders{})
}

func GetAllRoles() []Roles {
	var roles []Roles
	e := db.Find(&roles)
	// if e != nil {
	// 	sendErr(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	//e = json.NewEncoder(w).Encode(products)
	if e != nil {
		fmt.Println("Error in  getting all Roles")
	}
	return roles
}

func GetRoleByID(roleId string) Roles {
	var roles Roles
	e := db.First(&roles, roleId)
	if e != nil {
		fmt.Println("Error finding the Role")
	}
	return roles
}

func AddRole(roles *Roles) *Roles {

	e := db.Create(&roles)
	// fmt.Println("Hello")
	if e != nil {
		fmt.Println("Error creating a new Role")
	}
	return roles
}

func UpdateRole(roleId string, updateRole *Roles) Roles {
	var roles Roles
	db.First(&roles, roleId)
	//update role functionality is not working
	//db.Model(&roles).Where("Role_id=?", roleId).Updates(&UpdateRole)
	return roles
}

func DeleteRole(roleId string) Roles {
	var roles Roles
	db.Delete(&roles, roleId)
	// json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
	return roles
}

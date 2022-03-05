package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Staff struct {
	// gorm.Model
	STAFFS_ID  uint   `json:"staffs_id"`
	FIRST_NAME string `json:"first_name"`
	LAST_NAME  string `json:"last_name"`
	ADDRESS    string `json:"address"`
	PHONE      uint   `json:"phone"`
	EMAIL      string `json:"email"`
	USERNAME   string `json:"username"`
	PASSWORD   string `json:"password"`
	//ROLE_ID		uint 	`json:"role_id"`
	//role_id is foreign key...so for time being we are not manipulating it
}

func getAllStaff(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var staff []Staff
	e := db.Find(&staff).Error
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	e = json.NewEncoder(w).Encode(staff)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func getStaffByID(w http.ResponseWriter, r *http.Request) {
	var staff Staff
	queryParams := mux.Vars(r)
	db.Raw("select STAFFS_ID, FIRST_NAME, LAST_NAME, ADDRESS, PHONE, EMAIL, USERNAME, PASSWORD, ROLE_ID From STAFFS where USERNAME=?", queryParams["staffId"]).Scan(&staff)
	// data,_:= json.Marshal(&order)
	// fmt.Fprint(w, string(data))

	fmt.Println(queryParams["staffId"])
	// db.First(&staff, queryParams["staffId"])
	e := json.NewEncoder(w).Encode(staff)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func addStaff(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var staff Staff
	json.NewDecoder(r.Body).Decode(&staff)
	db.Create(&staff)
	json.NewEncoder(w).Encode(staff)
	// fmt.Println("Hello")
}

func updateStaff(w http.ResponseWriter, r *http.Request) {
	var staff Staff
	var updatedStaff Staff
	queryParams := mux.Vars(r)
	cusId := queryParams["supId"]
	db.First(&staff, cusId)
	json.NewDecoder(r.Body).Decode(&updatedStaff)
	db.Model(&staff).Where("customer_id=?", cusId).Updates(&updatedStaff)
	json.NewEncoder(w).Encode(&staff)

}

func deleteStaff(w http.ResponseWriter, r *http.Request) {
	var staff Staff
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["staffId"])
	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
	db.Delete(&staff, queryParams["staffId"])
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

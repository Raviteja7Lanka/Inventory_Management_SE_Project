package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Staffs struct {
	// gorm.Model
	STAFF_ID  string   `json:"staff_id"`
	FIRST_NAME string `json:"first_name"`
	LAST_NAME  string `json:"last_name"`
	ADDRESS    string `json:"address"`
	PHONE      uint   `json:"phone"`
	EMAIL      string `json:"email"`
	USERNAME   string `json:"username"`
	PASSWORD   string `json:"password"`
	ROLE_ID string `json:"role_id"`
}

func getAllStaff(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var staff []Staffs
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
	var staff Staffs
	queryParams := mux.Vars(r)
	db.Raw("select STAFF_ID, FIRST_NAME, LAST_NAME, ADDRESS, PHONE, EMAIL, USERNAME, PASSWORD, ROLE_ID From STAFFS where USERNAME=?", queryParams["staffId"]).Scan(&staff)
	fmt.Println(queryParams["staffId"])
	e := json.NewEncoder(w).Encode(staff)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func addStaff(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var staff Staffs
	json.NewDecoder(r.Body).Decode(&staff)
	e:=db.Create(&staff).Error
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
	}	
	json.NewEncoder(w).Encode(staff)
}

func updateStaff(w http.ResponseWriter, r *http.Request) {
	var staff Staffs
	var updatedStaffs Staffs
	queryParams := mux.Vars(r)
	cusId := queryParams["staffId"]
	db.First(&staff, cusId)
	json.NewDecoder(r.Body).Decode(&updatedStaffs)
	db.Model(&staff).Where("staff_id=?", cusId).Updates(&updatedStaffs)
	json.NewEncoder(w).Encode(&staff)

}

func deleteStaff(w http.ResponseWriter, r *http.Request) {
	var staff Staffs
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["staffId"])
	db.Delete(&staff, queryParams["staffId"])
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

package controllers

import (
	"apis/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// var db *gorm.DB

// func sendErr(w http.ResponseWriter, code int, message string) {
// 	resp, _ := json.Marshal(map[string]string{"error": message})
// 	http.Error(w, string(resp), code)
// }
// func GetAllStaff(w http.ResponseWriter, r *http.Request) {
// 	// w.Header().Set("Content-Type", "application/json")
// 	staff := models.GetAllStaff()
// 	res, e := json.Marshal(staff)
// 	if e != nil {
// 		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
// 		http.Error(w, string(resp), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func GetStaffByID(w http.ResponseWriter, r *http.Request) {

// 	queryParams := mux.Vars(r)

// 	staff := models.GetStaffByID(queryParams["staffId"])
// 	res, e := json.Marshal(staff)
// 	if e != nil {
// 		http.Error(w, string(e.Error()), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func AddStaff(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	staff := &models.Staff{}
// 	json.NewDecoder(r.Body).Decode(&staff)
// 	newCust := models.AddStaff(staff)
// 	res, e := json.Marshal(newCust)
// 	if e != nil {
// 		http.Error(w, string(e.Error()), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// 	// fmt.Println("Hello")
// }

// func UpdateStaff(w http.ResponseWriter, r *http.Request) {

// 	updatedStaff := &models.Staff{}
// 	queryParams := mux.Vars(r)
// 	staffId := queryParams["staffId"]
// 	json.NewDecoder(r.Body).Decode(&updatedStaff)
// 	staff := models.UpdateStaff(staffId, updatedStaff)
// 	res, e := json.Marshal(staff)
// 	if e != nil {
// 		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
// 		http.Error(w, string(resp), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }

// func DeleteStaff(w http.ResponseWriter, r *http.Request) {

// 	queryParams := mux.Vars(r)
// 	staff := models.DeleteStaff(queryParams["staffId"])
// 	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
// 	res, e := json.Marshal(staff)
// 	if e != nil {
// 		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
// 		http.Error(w, string(resp), http.StatusInternalServerError)
// 	}
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }

func GetAllStaff(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var staff []models.Staff
	e := db.Find(&staff).Error
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
		return
	}
	e = json.NewEncoder(w).Encode(staff)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
	}
}

func GetStaffByID(w http.ResponseWriter, r *http.Request) {
	var staff models.Staff
	queryParams := mux.Vars(r)
	db.Raw("select STAFFS_ID, FIRST_NAME, LAST_NAME, ADDRESS, PHONE, EMAIL, USERNAME, PASSWORD From STAFFS where EMAIL=?", queryParams["staffId"]).Scan(&staff)
	fmt.Println(queryParams["staffId"])
	e := json.NewEncoder(w).Encode(staff)
	if e != nil {
		sendErr(w, http.StatusInternalServerError, e.Error())
	}
}

func AddStaff(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var staff models.Staff
	json.NewDecoder(r.Body).Decode(&staff)
	db.Create(&staff)
	json.NewEncoder(w).Encode(staff)
}

func UpdateStaff(w http.ResponseWriter, r *http.Request) {
	var staff models.Staff
	var updatedStaff models.Staff
	queryParams := mux.Vars(r)
	cusId := queryParams["supId"]
	db.First(&staff, cusId)
	json.NewDecoder(r.Body).Decode(&updatedStaff)
	db.Model(&staff).Where("customer_id=?", cusId).Updates(&updatedStaff)
	json.NewEncoder(w).Encode(&staff)

}

func DeleteStaff(w http.ResponseWriter, r *http.Request) {
	var staff models.Staff
	queryParams := mux.Vars(r)
	fmt.Println(queryParams["staffId"])
	db.Delete(&staff, queryParams["staffId"])
	json.NewEncoder(w).Encode("{Status:200, Message: Deletion successful}")
}

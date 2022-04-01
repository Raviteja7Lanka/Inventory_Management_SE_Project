package controllers

import (
	"apis/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// var db *gorm.DB

func GetAllStaff(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	staff := models.GetAllStaff()
	res, e := json.Marshal(staff)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetStaffByID(w http.ResponseWriter, r *http.Request) {

	queryParams := mux.Vars(r)

	staff := models.GetStaffByID(queryParams["staffId"])
	res, e := json.Marshal(staff)
	if e != nil {
		http.Error(w, string(e.Error()), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AddStaff(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	staff := &models.Staff{}
	json.NewDecoder(r.Body).Decode(&staff)
	newCust := models.AddStaff(staff)
	res, e := json.Marshal(newCust)
	if e != nil {
		http.Error(w, string(e.Error()), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// fmt.Println("Hello")
}

func UpdateStaff(w http.ResponseWriter, r *http.Request) {

	updatedStaff := &models.Staff{}
	queryParams := mux.Vars(r)
	staffId := queryParams["staffId"]
	json.NewDecoder(r.Body).Decode(&updatedStaff)
	staff := models.UpdateStaff(staffId, updatedStaff)
	res, e := json.Marshal(staff)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteStaff(w http.ResponseWriter, r *http.Request) {

	queryParams := mux.Vars(r)
	staff := models.DeleteStaff(queryParams["staffId"])
	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
	res, e := json.Marshal(staff)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

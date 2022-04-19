package controllers

import (
	"apis/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// var db *gorm.DB

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	category := models.GetAllCategories()
	res, e := json.Marshal(category)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCategoriesByID(w http.ResponseWriter, r *http.Request) {
	// db.Raw("select id, date_of_order, order_details_id, customer_id, supplier_id, status from Orders where id=?",1).Scan(&order)
	// data,_:= json.Marshal(&order)
	// fmt.Fprint(w, string(data))
	queryParams := mux.Vars(r)
	// fmt.Println(queryParams["cusId"])
	// db.First(&customers, queryParams["cusId"])
	category := models.GetCategoriesByID(queryParams["catId"])
	res, e := json.Marshal(category)
	if e != nil {
		http.Error(w, string(e.Error()), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AddCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	category := &models.Category{}
	json.NewDecoder(r.Body).Decode(&category)
	newCat := models.AddCategories(category)
	res, e := json.Marshal(newCat)
	if e != nil {
		http.Error(w, string(e.Error()), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// fmt.Println("Hello")
}

func UpdateCategories(w http.ResponseWriter, r *http.Request) {

	updatedCategory := &models.Category{}
	queryParams := mux.Vars(r)
	catId := queryParams["catId"]
	json.NewDecoder(r.Body).Decode(&updatedCategory)
	category := models.UpdateCategories(catId, updatedCategory)
	res, e := json.Marshal(category)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteCategories(w http.ResponseWriter, r *http.Request) {

	queryParams := mux.Vars(r)
	category := models.DeleteCategories(queryParams["catId"])
	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
	res, e := json.Marshal(category)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

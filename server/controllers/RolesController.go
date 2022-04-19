package controllers

import (
	"apis/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// var db *gorm.DB

func GetAllRoles(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	roles := models.GetAllRoles()
	res, e := json.Marshal(roles)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetRoleByID(w http.ResponseWriter, r *http.Request) {
	// db.Raw("select id, date_of_order, order_details_id, customer_id, supplier_id, status from Orders where id=?",1).Scan(&order)
	// data,_:= json.Marshal(&order)
	// fmt.Fprint(w, string(data))
	queryParams := mux.Vars(r)
	// fmt.Println(queryParams["cusId"])
	// db.First(&customers, queryParams["cusId"])
	roles := models.GetRoleByID(queryParams["roleId"])
	res, e := json.Marshal(roles)
	if e != nil {
		http.Error(w, string(e.Error()), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AddRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	roles := &models.Roles{}
	json.NewDecoder(r.Body).Decode(&roles)
	newCust := models.AddRole(roles)
	res, e := json.Marshal(newCust)
	if e != nil {
		http.Error(w, string(e.Error()), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// fmt.Println("Hello")
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {

	// updatedRole := &models.Products{}
	// queryParams := mux.Vars(r)
	// // roleId := queryParams["roleId"]
	// json.NewDecoder(r.Body).Decode(&updatedRole)
	//roles := models.UpdateRole(roleId, updatedRole)
	//res, e := json.Marshal(roles)
	//if e != nil {
	//	resp, _ := json.Marshal(map[string]string{"error": e.Error()})
	//	http.Error(w, string(resp), http.StatusInternalServerError)
	//}
	// w.Header().Set("Content-Type", "pkglication/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)

}

func DeleteRole(w http.ResponseWriter, r *http.Request) {

	queryParams := mux.Vars(r)
	roles := models.DeleteRole(queryParams["roleId"])
	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
	res, e := json.Marshal(roles)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

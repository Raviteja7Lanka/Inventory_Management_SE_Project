package controllers

import (
	"apis/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// var db *gorm.DB

func GetAllWarehouses(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	warehouses := models.GetAllWarehouses()
	res, e := json.Marshal(warehouses)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetWarehouseByID(w http.ResponseWriter, r *http.Request) {
	// db.Raw("select id, date_of_order, order_details_id, customer_id, supplier_id, status from Orders where id=?",1).Scan(&order)
	// data,_:= json.Marshal(&order)
	// fmt.Fprint(w, string(data))
	queryParams := mux.Vars(r)
	// fmt.Println(queryParams["cusId"])
	// db.First(&customers, queryParams["cusId"])
	warehouses := models.GetWarehouseByID(queryParams["wareId"])
	res, e := json.Marshal(warehouses)
	if e != nil {
		http.Error(w, string(e.Error()), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AddWarehouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	warehouses := &models.Warehouses{}
	json.NewDecoder(r.Body).Decode(&warehouses)
	newWare := models.AddWarehouse(warehouses)
	res, e := json.Marshal(newWare)
	if e != nil {
		http.Error(w, string(e.Error()), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// fmt.Println("Hello")
}

func UpdateWarehouse(w http.ResponseWriter, r *http.Request) {

	updatedWarehouse := &models.Warehouses{}
	queryParams := mux.Vars(r)
	wareId := queryParams["wareId"]
	json.NewDecoder(r.Body).Decode(&updatedWarehouse)
	warehouses := models.UpdateWarehouse(wareId, updatedWarehouse)
	res, e := json.Marshal(warehouses)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteWarehouse(w http.ResponseWriter, r *http.Request) {

	queryParams := mux.Vars(r)
	warehouses := models.DeleteWarehouse(queryParams["wareId"])
	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
	res, e := json.Marshal(warehouses)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

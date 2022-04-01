package controllers

import (
	"apis/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// var db *gorm.DB

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	products := models.GetAllProducts()
	res, e := json.Marshal(products)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	// db.Raw("select id, date_of_order, order_details_id, customer_id, supplier_id, status from Orders where id=?",1).Scan(&order)
	// data,_:= json.Marshal(&order)
	// fmt.Fprint(w, string(data))
	queryParams := mux.Vars(r)
	// fmt.Println(queryParams["cusId"])
	// db.First(&customers, queryParams["cusId"])
	products := models.GetProductByID(queryParams["prodId"])
	res, e := json.Marshal(products)
	if e != nil {
		http.Error(w, string(e.Error()), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products := &models.Products{}
	json.NewDecoder(r.Body).Decode(&products)
	newCust := models.AddProduct(products)
	res, e := json.Marshal(newCust)
	if e != nil {
		http.Error(w, string(e.Error()), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// fmt.Println("Hello")
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	updatedProduct := &models.Products{}
	queryParams := mux.Vars(r)
	prodId := queryParams["prodId"]
	json.NewDecoder(r.Body).Decode(&updatedProduct)
	products := models.UpdateProduct(prodId, updatedProduct)
	res, e := json.Marshal(products)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	queryParams := mux.Vars(r)
	products := models.DeleteProduct(queryParams["prodId"])
	//db.Raw("delete from Orders where order_id=?", queryParams["ordId"])
	res, e := json.Marshal(products)
	if e != nil {
		resp, _ := json.Marshal(map[string]string{"error": e.Error()})
		http.Error(w, string(resp), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
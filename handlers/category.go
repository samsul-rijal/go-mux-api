package handlers

import (
	"encoding/json"
	"go-mux-api/models"
	"go-mux-api/pkg/mysql"
	"io/ioutil"
	"net/http"
)

func CategoryCreate(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var category models.Category
	json.Unmarshal(payloads, &category)

	errCreateCategory := mysql.DB.Create(&category).Error

	if errCreateCategory != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Result{Code: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": errCreateCategory.Error()}}
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
	response := Result{Code: http.StatusOK, Message: "success", Data: map[string]interface{}{"product": category}}
	json.NewEncoder(w).Encode(response)
}

func CategoryGetAll(w http.ResponseWriter, r *http.Request) {
	category := []models.Category{}
	mysql.DB.Find(&category)

	res := Result{Code: http.StatusOK, Message: "Success get user", Data: category}
	results, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)

}

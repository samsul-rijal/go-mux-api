package handlers

import (
	"encoding/json"
	"go-mux-api/models"
	"go-mux-api/pkg/mysql"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func UserGetAll(w http.ResponseWriter, r *http.Request) {
	users := []models.UserResponseWithProduct{}
	mysql.DB.Preload("Products").Preload("Profile").Find(&users)

	res := Result{Code: http.StatusOK, Data: users, Message: "Success get user"}
	results, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)

}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var user models.User
	json.Unmarshal(payloads, &user)

	mysql.DB.Create(&user)
	res := Result{Code: 201, Data: user, Message: "Create user success"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func UserGetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	var user models.User

	err := mysql.DB.Preload("Products").First(&user, userId).Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Result{Code: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := Result{Code: http.StatusOK, Message: "success", Data: map[string]interface{}{"user": user}}
	json.NewEncoder(w).Encode(response)

}

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
	users := []models.User{}
	mysql.DB.Preload("Products").Preload("Profile").Find(&users)

	// res := Result{Code: http.StatusOK, Data: users, Message: "Success get user"}
	// results, err := json.Marshal(res)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// w.Write(results)
	json.NewEncoder(w).Encode(users)

}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var user models.User
	json.Unmarshal(body, &user)

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

	err := mysql.DB.Debug().First(&user, userId).Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Result{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := Result{Code: http.StatusOK, Message: "success", Data: map[string]interface{}{"user": user}}
	json.NewEncoder(w).Encode(response)

}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	var user models.User

	err := mysql.DB.First(&user, userId).Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := Result{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	body := json.NewDecoder(r.Body)
	if err := body.Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := Result{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := mysql.DB.Debug().Where("id = ?", userId).Updates(&user).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := Result{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := Result{Code: http.StatusOK, Message: "success", Data: user}
	json.NewEncoder(w).Encode(response)
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	var user models.User

	err := mysql.DB.First(&user, userId).Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := Result{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := mysql.DB.Debug().Unscoped().Delete(&user, userId).Error; err != nil {
		w.Header().Set("Content-Type", "application/json")
		response := Result{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := Result{Code: http.StatusOK, Message: "Delete user success"}
	json.NewEncoder(w).Encode(response)
}

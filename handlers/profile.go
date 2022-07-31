package handlers

import (
	"encoding/json"
	"go-mux-api/models"
	"go-mux-api/pkg/mysql"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func ProfileGet(w http.ResponseWriter, r *http.Request) {
	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	var profile models.Profile
	err := mysql.DB.First(&profile, userId).Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Result{Code: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := Result{Code: http.StatusOK, Message: "success", Data: map[string]interface{}{"user": profile}}
	json.NewEncoder(w).Encode(response)

}

package handlers

import (
	"encoding/json"
	"fmt"
	"go-mux-api/models"
	jwtToken "go-mux-api/pkg/jwt"
	"go-mux-api/pkg/mysql"
	"go-mux-api/pkg/password"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Register(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var user models.User
	json.Unmarshal(payloads, &user)

	hashedPassword, err := password.HashingPassword(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Result{Code: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
	}

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}

	errCreateUser := mysql.DB.Create(&newUser).Error
	if errCreateUser != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Result{Code: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
	}

	// create profile user
	newProfile := models.Profile{
		UserID: newUser.ID,
	}
	errCreateProfile := mysql.DB.Create(&newProfile).Error

	if errCreateProfile != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Result{Code: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
	response := Result{Code: http.StatusOK, Message: "success", Data: map[string]interface{}{"user": newUser}}
	json.NewEncoder(w).Encode(response)

}

func Login(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var user models.User
	json.Unmarshal(payloads, &user)

	newUser := models.User{
		Email:    user.Email,
		Password: user.Password,
	}

	// Check email
	err := mysql.DB.First(&user, "email = ?", newUser.Email).Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := Result{Code: http.StatusBadRequest, Message: "wrong email or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check password
	isValid := password.CheckPasswordHash(newUser.Password, user.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := Result{Code: http.StatusBadRequest, Message: "wrong email or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	// claims["email"] = user.Email
	// claims["status"] = user.Status
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 jam expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
		return
	}

	result := models.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	response := Result{Code: http.StatusOK, Message: "success", Data: map[string]interface{}{"user": result, "token": token}}
	json.NewEncoder(w).Encode(response)

}

package handlers

import (
	"encoding/json"
	"go-mux-api/models"
	"go-mux-api/pkg/mysql"
	"io/ioutil"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func ProductCreate(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var product models.Product
	json.Unmarshal(payloads, &product)

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	newProduct := models.Product{
		Name:   product.Name,
		Desc:   product.Desc,
		Price:  product.Price,
		Image:  product.Image,
		Qty:    product.Qty,
		UserID: userId,
	}

	errCreateProduct := mysql.DB.Debug().Create(&newProduct).Error

	if errCreateProduct != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Result{Code: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": errCreateProduct.Error()}}
		json.NewEncoder(w).Encode(response)
	}

	w.Header().Set("Content-Type", "application/json")
	response := Result{Code: http.StatusOK, Message: "success", Data: map[string]interface{}{"product": newProduct}}
	json.NewEncoder(w).Encode(response)
}

func ProductGetAll(w http.ResponseWriter, r *http.Request) {
	products := []models.Product{}
	mysql.DB.Find(&products)
	res := Result{Code: http.StatusOK, Data: products, Message: "Success get user"}

	results, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)

}

package handlers

import (
	"encoding/json"
	"go-mux-api/models"
	"go-mux-api/pkg/mysql"
	"io/ioutil"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func TransactionGetAll(w http.ResponseWriter, r *http.Request) {
	products := []models.Transaction{}
	mysql.DB.Preload("Product").Preload("Product.User").Preload("Buyer").Preload("Seller").Find(&products)

	res := Result{Code: http.StatusOK, Data: products, Message: "Success get transaction"}
	results, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)

}

func TransactionCreate(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var transaction models.Transaction
	json.Unmarshal(payloads, &transaction)

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	newTransaction := models.Transaction{
		ProductID: transaction.ProductID,
		BuyerID:   userId,
		SellerID:  transaction.SellerID,
		Price:     transaction.Price,
		Status:    "pending",
	}

	errCreateProduct := mysql.DB.Create(&newTransaction).Error

	if errCreateProduct != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Result{Code: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": errCreateProduct.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := Result{Code: http.StatusOK, Message: "success", Data: map[string]interface{}{"product": newTransaction}}
	json.NewEncoder(w).Encode(response)
}

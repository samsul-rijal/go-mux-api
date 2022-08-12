package handlers

import (
	"encoding/json"
	"fmt"
	"go-mux-api/models"
	"go-mux-api/pkg/mysql"
	"io/ioutil"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/midtrans/midtrans-go"
	snap "github.com/midtrans/midtrans-go/snap"
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

	var user models.User
	err := mysql.DB.Debug().First(&user, newTransaction.BuyerID).Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Result{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// 1. Initiate Snap client
	var s snap.Client
	s.New("SB-Mid-server-9-9vMu3s_szn7q6pU74APTvH", midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "ORDER-ID-12345",
			GrossAmt: int64(newTransaction.Price),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)
	fmt.Println(snapResp)

	w.Header().Set("Content-Type", "application/json")
	response := Result{Code: http.StatusOK, Message: "success", Data: map[string]interface{}{"product": newTransaction, "url": snapResp}}
	json.NewEncoder(w).Encode(response)
}

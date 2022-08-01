package handlers

import (
	"encoding/json"
	"go-mux-api/models"
	"go-mux-api/pkg/mysql"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

func ProductCreate(w http.ResponseWriter, r *http.Request) {
	// payloads, _ := ioutil.ReadAll(r.Body)

	var product models.Product
	// json.Unmarshal(payloads, &product)

	// get data filename middleware
	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	newProduct := models.Product{
		Name:   product.Name,
		Desc:   product.Desc,
		Price:  product.Price,
		Image:  filename,
		Qty:    product.Qty,
		UserID: userId,
	}

	errCreateProduct := mysql.DB.Create(&newProduct).Error

	// create data many2many
	for _, categoryId := range product.CategoryID {
		productCategory := models.ProductCategory{
			ProductID:  newProduct.ID,
			CategoryID: categoryId,
		}
		mysql.DB.Debug().Create(&productCategory)
	}

	if errCreateProduct != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Result{Code: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": errCreateProduct.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := Result{Code: http.StatusOK, Message: "success", Data: map[string]interface{}{"product": newProduct}}
	json.NewEncoder(w).Encode(response)
}

func ProductGetAll(w http.ResponseWriter, r *http.Request) {
	products := []models.ProductResponseWithCategory{}
	mysql.DB.Preload("User").Preload("Category").Find(&products)

	log.Println(products)

	res := Result{Code: http.StatusOK, Data: products, Message: "Success get user"}
	results, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func ProductGetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productId := params["id"]
	var product models.Product

	err := mysql.DB.Preload("User").First(&product, productId).Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := Result{Code: http.StatusBadRequest, Message: "product not found"}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := Result{Code: http.StatusOK, Message: "success", Data: map[string]interface{}{"product": product}}
	json.NewEncoder(w).Encode(response)

}

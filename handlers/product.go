package handlers

import (
	"encoding/json"
	"go-mux-api/models"
	"go-mux-api/pkg/mysql"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

func ProductCreate(w http.ResponseWriter, r *http.Request) {

	// payloads, _ := ioutil.ReadAll(r.Body)

	var product models.Product
	// json.Unmarshal(payloads, &product)

	// fmt.Fprintln(w, "Name : ", r.Form.Get("name"))

	// get data filename middleware
	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// handle request form data
	name := r.FormValue("name")
	desc := r.FormValue("desc")
	price, _ := strconv.Atoi(r.FormValue("price"))
	qty, _ := strconv.Atoi(r.FormValue("qty"))

	newProduct := models.Product{
		Name:   name,
		Desc:   desc,
		Price:  price,
		Image:  filename,
		Qty:    qty,
		UserID: userId,
	}

	err := mysql.DB.Create(&newProduct).Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := Result{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// create data many2many
	for _, categoryId := range product.CategoryID {
		productCategory := models.ProductCategory{
			ProductID:  newProduct.ID,
			CategoryID: categoryId,
		}
		mysql.DB.Debug().Create(&productCategory)
	}

	w.Header().Set("Content-Type", "application/json")
	response := Result{Code: http.StatusOK, Message: "success", Data: map[string]interface{}{"product": newProduct}}
	json.NewEncoder(w).Encode(response)
}

func ProductGetAll(w http.ResponseWriter, r *http.Request) {
	products := []models.ProductResponseWithCategory{}
	mysql.DB.Preload("User").Preload("Category").Find(&products)

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

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
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

	// get data filename middleware
	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// handle request form data
	name := r.FormValue("name")
	desc := r.FormValue("desc")
	price, _ := strconv.Atoi(r.FormValue("price"))
	qty, _ := strconv.Atoi(r.FormValue("qty"))

	newProduct := models.Product{
		Name:   name,
		Desc:   desc,
		Price:  price,
		Image:  filename,
		Qty:    qty,
		UserID: userId,
	}

	errUpdate := mysql.DB.Debug().Where("id = ?", productId).Updates(&newProduct).Error
	if errUpdate != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := Result{Code: http.StatusBadRequest, Message: errUpdate.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// create data many2many
	for _, categoryId := range product.CategoryID {
		productCategory := models.ProductCategory{
			ProductID:  newProduct.ID,
			CategoryID: categoryId,
		}
		mysql.DB.Debug().Create(&productCategory)
	}

	w.Header().Set("Content-Type", "application/json")
	response := Result{Code: http.StatusOK, Message: "success", Data: map[string]interface{}{"product": newProduct}}
	json.NewEncoder(w).Encode(response)

}

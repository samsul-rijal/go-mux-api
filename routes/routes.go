package routes

import (
	"encoding/json"
	"go-mux-api/handlers"
	"go-mux-api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Hello World")
}

func RouteInit(r *mux.Router) {

	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	r.HandleFunc("/users", handlers.UserGetAll).Methods("GET")
	r.HandleFunc("/user/{id}", handlers.UserGetById).Methods("GET")
	r.HandleFunc("/user", handlers.UserCreate).Methods("POST")
	r.HandleFunc("/user/{id}", handlers.UserUpdate).Methods("PUT")
	r.HandleFunc("/user/{id}", handlers.UserDelete).Methods("DELETE")

	r.HandleFunc("/profile", middleware.Auth(handlers.ProfileGet)).Methods("GET")

	// r.HandleFunc("/product", middleware.Auth(handlers.ProductCreate)).Methods("POST")
	r.HandleFunc("/product", middleware.Auth(middleware.UploadFile(handlers.ProductCreate))).Methods("POST")
	r.HandleFunc("/products", middleware.Auth(handlers.ProductGetAll)).Methods("GET")
	r.HandleFunc("/product/{id}", middleware.Auth(handlers.ProductGetById)).Methods("GET")
	r.HandleFunc("/product/{id}", middleware.Auth(middleware.UploadFile(handlers.UpdateProduct))).Methods("PUT")

	r.HandleFunc("/categories", middleware.Auth(handlers.CategoryGetAll)).Methods("GET")
	r.HandleFunc("/category", middleware.Auth(handlers.CategoryCreate)).Methods("POST")

	r.HandleFunc("/transactions", middleware.Auth(handlers.TransactionGetAll)).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(handlers.TransactionCreate)).Methods("POST")

	r.HandleFunc("/upload", middleware.Auth(middleware.UploadFile(handlers.UpdateProduct))).Methods("POST")
	r.HandleFunc("/hello", Hello).Methods("GET")
}

package routes

import (
	"fmt"
	"go-mux-api/handlers"
	"go-mux-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RouteInit() {
	r := mux.NewRouter()

	r.HandleFunc("/user", middleware.Auth(handlers.UserGetAll)).Methods("GET")
	r.HandleFunc("/user/{id}", handlers.UserGetById).Methods("GET")
	r.HandleFunc("/user", handlers.UserCreate).Methods("POST")
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	r.HandleFunc("/product", middleware.Auth(handlers.ProductCreate)).Methods("POST")
	r.HandleFunc("/products", middleware.Auth(handlers.ProductGetAll)).Methods("GET")

	fmt.Println("server running localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}

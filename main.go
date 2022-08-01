package main

import (
	"fmt"
	"go-mux-api/database"
	"go-mux-api/pkg/mysql"
	"go-mux-api/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	mysql.DatabaseInit()
	database.RunMigration()

	// userRepository := repository.RepositoryUser(mysql.DB)

	r := mux.NewRouter()
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}

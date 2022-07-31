package main

import (
	"go-mux-api/database"
	"go-mux-api/pkg/mysql"
	"go-mux-api/routes"
)

func main() {

	mysql.DatabaseInit()
	database.RunMigration()

	routes.RouteInit()
}

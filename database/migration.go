package database

import (
	"fmt"
	"go-mux-api/models"
	"go-mux-api/pkg/mysql"
)

func RunMigration() {
	// database.DB.AutoMigrate(&entity.User{}, &next-entity)
	err := mysql.DB.AutoMigrate(&models.User{}, &models.Product{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}

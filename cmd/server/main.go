package main

import (
	"github.com/claytonssmint/devfullcycle/goexpert/APIS/configs"
	"github.com/claytonssmint/devfullcycle/goexpert/APIS/internal/entity"
	"github.com/claytonssmint/devfullcycle/goexpert/APIS/internal/infra/database"
	"github.com/claytonssmint/devfullcycle/goexpert/APIS/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	producthandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", producthandler.CreateProduct)
	http.ListenAndServe(":8000", nil)

}

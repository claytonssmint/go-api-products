package main

import (
	"github.com/claytonssmint/devfullcycle/goexpert/APIS/configs"
	"github.com/claytonssmint/devfullcycle/goexpert/APIS/internal/entity"
	"github.com/claytonssmint/devfullcycle/goexpert/APIS/internal/infra/database"
	"github.com/claytonssmint/devfullcycle/goexpert/APIS/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", producthandler.CreateProduct)

	http.ListenAndServe(":8000", r)
}

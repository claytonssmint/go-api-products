package handlers

import (
	"encoding/json"
	"github.com/claytonssmint/devfullcycle/goexpert/APIS/internal/dto"
	"github.com/claytonssmint/devfullcycle/goexpert/APIS/internal/entity"
	"github.com/claytonssmint/devfullcycle/goexpert/APIS/internal/infra/database"
	"net/http"
)

type Producthandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *Producthandler {
	return &Producthandler{
		ProductDB: db,
	}
}

func (h *Producthandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

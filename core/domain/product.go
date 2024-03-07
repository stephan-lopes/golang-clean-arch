package domain

import (
	"net/http"

	"github.com/stephan-lopes/golang-clean-arch/core/dto"
)

type Product struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

type ProductService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Fetch(w http.ResponseWriter, r *http.Request)
}

type ProductUseCase interface {
	Create(productRequest *dto.CreateProductRequest) (*Product, error)
	Fetch(paginationRequest *dto.PaginationRequestParams) (*Pagination[[]Product], error)
}

type ProductRepository interface {
	Create(productRequest *dto.CreateProductRequest) (*Product, error)
	Fetch(paginationRequest *dto.PaginationRequestParams) (*Pagination[[]Product], error)
}

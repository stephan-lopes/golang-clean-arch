package domain

import (
	"net/http"

	"github.com/stephan-lopes/golang-clean-arch/core/dto"
)

type Product struct {
	ID          int32   `json:"id" example:"1"`
	Name        string  `json:"name" example:"Mesa"`
	Price       float32 `json:"price" example:"200.00"`
	Description string  `json:"description" example:"Uma mesa, como outra qualquer"`
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

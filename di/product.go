package di

import (
	"github.com/stephan-lopes/golang-clean-arch/adapter/http/productservice"
	"github.com/stephan-lopes/golang-clean-arch/adapter/postgres"
	"github.com/stephan-lopes/golang-clean-arch/adapter/postgres/productrepository"
	"github.com/stephan-lopes/golang-clean-arch/core/domain"
	"github.com/stephan-lopes/golang-clean-arch/core/domain/usecase/productusecase"
)

func ConfigProductDI(conn postgres.PoolInterface) domain.ProductService {
	productRepository := productrepository.New(conn)
	productUseCase := productusecase.New(productRepository)
	productService := productservice.New(productUseCase)

	return productService
}

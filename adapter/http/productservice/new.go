package productservice

import "github.com/stephan-lopes/golang-clean-arch/core/domain"

type service struct {
	usecase domain.ProductUseCase
}

func New(usecase domain.ProductUseCase) domain.ProductService {
	return &service{
		usecase: usecase,
	}
}

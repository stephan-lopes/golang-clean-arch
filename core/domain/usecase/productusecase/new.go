package productusecase

import "github.com/stephan-lopes/golang-clean-arch/core/domain"

type usecase struct {
	repository domain.ProductRepository
}

func New(repository domain.ProductRepository) domain.ProductUseCase {
	return &usecase{
		repository: repository,
	}
}

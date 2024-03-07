package productusecase

import (
	"github.com/stephan-lopes/golang-clean-arch/core/domain"
	"github.com/stephan-lopes/golang-clean-arch/core/dto"
)

func (usecase usecase) Fetch(paginationRequest *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Product], error) {
	products, err := usecase.repository.Fetch(paginationRequest)

	if err != nil {
		return nil, err
	}

	return products, nil
}

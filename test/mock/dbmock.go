package mock

import (
	"github.com/go-faker/faker/v4"
	"github.com/pashagolub/pgxmock"
	"github.com/stephan-lopes/golang-clean-arch/core/domain"
	"github.com/stephan-lopes/golang-clean-arch/core/dto"
)

func NewDBSetupCreate() ([]string, dto.CreateProductRequest, domain.Product, pgxmock.PgxPoolIface) {
	cols := []string{"id", "name", "price", "description"}
	fakeProductRequest := dto.CreateProductRequest{}
	fakeProductDBResponse := domain.Product{}
	faker.FakeData(&fakeProductRequest)
	faker.FakeData(&fakeProductDBResponse)

	mock, _ := pgxmock.NewPool()

	return cols, fakeProductRequest, fakeProductDBResponse, mock
}

func NewDBSetupFetch() ([]string, dto.PaginationRequestParams, domain.Product, pgxmock.PgxPoolIface) {
	cols := []string{"id", "name", "price", "description"}
	fakePaginationRequestParams := dto.PaginationRequestParams{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         nil,
		Descending:   nil,
		Search:       "",
	}
	fakeProductDBResponse := domain.Product{}
	faker.FakeData(&fakeProductDBResponse)

	mock, _ := pgxmock.NewPool()

	return cols, fakePaginationRequestParams, fakeProductDBResponse, mock
}

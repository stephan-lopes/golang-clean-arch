package productusecase_test

import (
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stephan-lopes/golang-clean-arch/core/domain"
	"github.com/stephan-lopes/golang-clean-arch/core/domain/usecase/productusecase"
	"github.com/stephan-lopes/golang-clean-arch/core/dto"
	"github.com/stephan-lopes/golang-clean-arch/test/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestFetch(t *testing.T) {
	fakePaginationRequestParams := dto.PaginationRequestParams{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         nil,
		Descending:   nil,
		Search:       "",
	}
	fakeDBProduct := domain.Product{}

	faker.FakeData(&fakeDBProduct)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProduckRepository := mock.NewMockProductRepository(mockCtrl)
	mockProduckRepository.EXPECT().Fetch(&fakePaginationRequestParams).Return(&domain.Pagination[[]domain.Product]{
		Items: []domain.Product{fakeDBProduct},
		Total: 1,
	}, nil)

	sut := productusecase.New(mockProduckRepository)
	products, err := sut.Fetch(&fakePaginationRequestParams)

	require.Nil(t, err)

	for _, product := range products.Items {
		require.Nil(t, err)
		require.NotEmpty(t, product.ID)
		require.Equal(t, product.Name, fakeDBProduct.Name)
		require.Equal(t, product.Price, fakeDBProduct.Price)
		require.Equal(t, product.Description, fakeDBProduct.Description)
	}
}

func TestFetch_Error(t *testing.T) {
	fakePaginationRequestParams := dto.PaginationRequestParams{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         nil,
		Descending:   nil,
		Search:       "",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProductRepository := mock.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().Fetch(&fakePaginationRequestParams).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := productusecase.New(mockProductRepository)
	product, err := sut.Fetch(&fakePaginationRequestParams)

	require.NotNil(t, err)
	require.Nil(t, product)
}

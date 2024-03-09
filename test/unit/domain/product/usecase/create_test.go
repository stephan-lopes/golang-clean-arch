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

func TestCreate(t *testing.T) {
	fakeRequestProduct := dto.CreateProductRequest{}
	fakeDBProduct := domain.Product{}
	faker.FakeData(&fakeRequestProduct)
	faker.FakeData(&fakeDBProduct)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mock.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().Create(&fakeRequestProduct).Return(&fakeDBProduct, nil)

	sut := productusecase.New(mockProductRepository)
	product, err := sut.Create(&fakeRequestProduct)

	require.Nil(t, err)
	require.NotEmpty(t, product.ID)
	require.Equal(t, product.Name, fakeDBProduct.Name)
	require.Equal(t, product.Price, fakeDBProduct.Price)
	require.Equal(t, product.Description, fakeDBProduct.Description)
}

func TestCreate_Error(t *testing.T) {
	fakeRequestProduct := dto.CreateProductRequest{}
	faker.FakeData(&fakeRequestProduct)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProduckRepository := mock.NewMockProductRepository(mockCtrl)
	mockProduckRepository.EXPECT().Create(&fakeRequestProduct).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := productusecase.New(mockProduckRepository)
	product, err := sut.Create(&fakeRequestProduct)

	require.NotNil(t, err)
	require.Nil(t, product)
}

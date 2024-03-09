package productservice__test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stephan-lopes/golang-clean-arch/adapter/http/productservice"
	"github.com/stephan-lopes/golang-clean-arch/core/domain"
	"github.com/stephan-lopes/golang-clean-arch/core/dto"
	"github.com/stephan-lopes/golang-clean-arch/test/mock"
	"go.uber.org/mock/gomock"
)

func TestFetch(t *testing.T) {
	fakePaginationRequestParams := dto.PaginationRequestParams{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         []string{""},
		Descending:   []string{""},
		Search:       "",
	}

	fakeProduct := domain.Product{}
	faker.FakeData(&fakeProduct)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProductUseCase := mock.NewMockProductUseCase(mockCtrl)
	mockProductUseCase.EXPECT().Fetch(&fakePaginationRequestParams).Return(&domain.Pagination[[]domain.Product]{
		Items: []domain.Product{fakeProduct},
		Total: 1,
	}, nil)

	sut := productservice.New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/product", nil)
	r.Header.Set("Content-Type", "application/json")
	queryStringParams := r.URL.Query()
	queryStringParams.Add("page", "1")
	queryStringParams.Add("itemsPerPage", "10")
	queryStringParams.Add("sort", "")
	queryStringParams.Add("descending", "")
	queryStringParams.Add("search", "")

	r.URL.RawQuery = queryStringParams.Encode()
	sut.Fetch(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("status code is not correct")
	}
}

func TestFetch_ProductError(t *testing.T) {
	fakePaginationRequestParams := dto.PaginationRequestParams{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         []string{""},
		Descending:   []string{""},
		Search:       "",
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductUseCase := mock.NewMockProductUseCase(mockCtrl)
	mockProductUseCase.EXPECT().Fetch(&fakePaginationRequestParams).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := productservice.New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/product", nil)
	r.Header.Set("Content-Type", "application/json")
	queryStringParams := r.URL.Query()
	queryStringParams.Add("page", "1")
	queryStringParams.Add("itemsPerPage", "10")
	queryStringParams.Add("sort", "")
	queryStringParams.Add("descending", "")
	queryStringParams.Add("search", "")
	r.URL.RawQuery = queryStringParams.Encode()
	sut.Fetch(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}

}

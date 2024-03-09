package productservice__test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stephan-lopes/golang-clean-arch/adapter/http/productservice"
	"github.com/stephan-lopes/golang-clean-arch/core/domain"
	"github.com/stephan-lopes/golang-clean-arch/core/dto"
	"github.com/stephan-lopes/golang-clean-arch/test/mock"
	"go.uber.org/mock/gomock"
)

func TestCreate(t *testing.T) {
	fakeProductRequest := dto.CreateProductRequest{}
	fakeProduct := domain.Product{}
	faker.FakeData(&fakeProductRequest)
	faker.FakeData(&fakeProduct)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProductUseCase := mock.NewMockProductUseCase(mockCtrl)
	mockProductUseCase.EXPECT().Create(&fakeProductRequest).Return(&fakeProduct, nil)

	payload, _ := json.Marshal(fakeProductRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")

	sut := productservice.New(mockProductUseCase)
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("status code is not correct")
	}
}

func TestCreate_ProductError(t *testing.T) {
	fakeProductRequest := dto.CreateProductRequest{}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUseCase := mock.NewMockProductUseCase(mockCtrl)
	mockUseCase.EXPECT().Create(&fakeProductRequest).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := productservice.New(mockUseCase)

	payload, _ := json.Marshal(fakeProductRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/product", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}

func TestCreate_JsonErrorFormater(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUseCase := mock.NewMockProductUseCase(mockCtrl)

	sut := productservice.New(mockUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader("{"))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}

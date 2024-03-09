package productservice

import (
	"encoding/json"
	"net/http"

	"github.com/stephan-lopes/golang-clean-arch/core/dto"
)

func (service service) Create(w http.ResponseWriter, r *http.Request) {
	productRequest, err := dto.FromJSONCreateProductRequest(r.Body)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	product, err := service.usecase.Create(productRequest)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(product)
}

package productservice

import (
	"encoding/json"
	"net/http"

	"github.com/stephan-lopes/golang-clean-arch/core/dto"
)

func (service service) Fetch(w http.ResponseWriter, r *http.Request) {
	paginationRequest, err := dto.FromValuePaginationRequestParams(r)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	products, err := service.usecase.Fetch(paginationRequest)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(products)
}

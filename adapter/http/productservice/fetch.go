package productservice

import (
	"encoding/json"
	"net/http"

	"github.com/stephan-lopes/golang-clean-arch/core/dto"
)

// @Summary Fetch products with server pagination
// @Description Fetch products with server pagination
// @Tags product
// @Accept  json
// @Produce  json
// @Param sort query string true "1,2"
// @Param descending query string true "true,false"
// @Param page query integer true "1"
// @Param itemsPerPage query integer true "10"
// @Param search query string false "1,2"
// @Success 200 {object} domain.Product
// @Router /product [get]
func (service service) Fetch(w http.ResponseWriter, r *http.Request) {
	paginationRequest, _ := dto.FromValuePaginationRequestParams(r)

	products, err := service.usecase.Fetch(paginationRequest)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(products)
}

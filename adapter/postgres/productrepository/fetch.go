package productrepository

import (
	"context"

	"github.com/booscaaa/go-paginate/paginate"
	"github.com/stephan-lopes/golang-clean-arch/core/domain"
	"github.com/stephan-lopes/golang-clean-arch/core/dto"
)

func (repository repository) Fetch(paginationRequest *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Product], error) {
	ctx := context.Background()
	products := []domain.Product{}
	total := int32(0)
	pager := paginate.Instance("")

	query, queryCount := pager.
		Query("SELECT * FROM product").
		Page(paginationRequest.Page).
		Desc(paginationRequest.Descending).
		Sort(paginationRequest.Sort).
		RowsPerPage(paginationRequest.ItemsPerPage).
		SearchBy(paginationRequest.Search, "name", "description").
		Select()

	rows, err := repository.db.Query(
		ctx,
		*query,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := domain.Product{}

		rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Description,
		)

		products = append(products, product)
	}

	err = repository.db.QueryRow(ctx, *queryCount).Scan(&total)
	if err != nil {
		return nil, err
	}

	return &domain.Pagination[[]domain.Product]{
		Items: products,
		Total: total,
	}, nil

}

package productrepository_test

import (
	"fmt"
	"testing"

	"github.com/pashagolub/pgxmock"
	"github.com/stephan-lopes/golang-clean-arch/adapter/postgres/productrepository"
	"github.com/stephan-lopes/golang-clean-arch/test/mock"
	"github.com/stretchr/testify/require"
)

func TestFetch(t *testing.T) {
	cols, fakePaginationRequestParams, fakeProductDBResponse, mockDB := mock.NewDBSetupFetch()
	defer mockDB.Close()

	mockDB.ExpectQuery("SELECT (.+) FROM product").
		WillReturnRows(pgxmock.NewRows(cols).AddRow(
			fakeProductDBResponse.ID,
			fakeProductDBResponse.Name,
			fakeProductDBResponse.Price,
			fakeProductDBResponse.Description,
		))

	mockDB.ExpectQuery("SELECT COUNT(.+) FROM product").
		WillReturnRows(pgxmock.NewRows([]string{"count"}).AddRow(int32(1)))

	sut := productrepository.New(mockDB)
	products, err := sut.Fetch(&fakePaginationRequestParams)

	require.Nil(t, err)

	if err := mockDB.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	for _, product := range products.Items {
		require.Nil(t, err)
		require.NotEmpty(t, product.ID)
		require.Equal(t, product.Name, fakeProductDBResponse.Name)
		require.Equal(t, product.Price, fakeProductDBResponse.Price)
		require.Equal(t, product.Description, fakeProductDBResponse.Description)
	}
}

func TestFetch_QueryError(t *testing.T) {
	_, fakePaginationRequestParams, _, mockDB := mock.NewDBSetupFetch()
	defer mockDB.Close()

	mockDB.ExpectQuery("SELECT (.+) FROM product").
		WillReturnError(fmt.Errorf("ANY QUERY ERROR"))

	sut := productrepository.New(mockDB)
	products, err := sut.Fetch(&fakePaginationRequestParams)

	require.NotNil(t, err)

	if err := mockDB.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, products)
}

func TestFetch_QueryCountError(t *testing.T) {
	cols, fakePaginationRequestParams, fakeProductDBResponse, mockDB := mock.NewDBSetupFetch()
	defer mockDB.Close()

	mockDB.ExpectQuery("SELECT (.+) FROM product").
		WillReturnRows(pgxmock.NewRows(cols).AddRow(
			fakeProductDBResponse.ID,
			fakeProductDBResponse.Name,
			fakeProductDBResponse.Price,
			fakeProductDBResponse.Description,
		))

	mockDB.ExpectQuery("SELECT COUNT(.+) FROM product").
		WillReturnError(fmt.Errorf("ANY QUERY COUNT ERROR"))

	sut := productrepository.New(mockDB)
	products, err := sut.Fetch(&fakePaginationRequestParams)

	require.NotNil(t, err)

	if err := mockDB.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	require.Nil(t, products)
}

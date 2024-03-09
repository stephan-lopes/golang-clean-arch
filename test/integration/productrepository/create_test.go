package productrepository_test

import (
	"fmt"
	"testing"

	"github.com/pashagolub/pgxmock"
	"github.com/stephan-lopes/golang-clean-arch/adapter/postgres/productrepository"
	"github.com/stephan-lopes/golang-clean-arch/test/mock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	cols, fakeProductRequest, fakeProductDBResponse, mockDB := mock.NewDBSetupCreate()
	defer mockDB.Close()

	mockDB.ExpectQuery("INSERT INTO product (.+)").WithArgs(
		fakeProductRequest.Name,
		fakeProductRequest.Price,
		fakeProductRequest.Description,
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeProductDBResponse.ID,
		fakeProductDBResponse.Name,
		fakeProductDBResponse.Price,
		fakeProductDBResponse.Description,
	))

	sut := productrepository.New(mockDB)
	product, err := sut.Create(&fakeProductRequest)

	if err := mockDB.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.NotEmpty(t, product.ID)
	require.Equal(t, product.Name, fakeProductDBResponse.Name)
	require.Equal(t, product.Price, fakeProductDBResponse.Price)
	require.Equal(t, product.Description, fakeProductDBResponse.Description)
}

func TestCreate_DBError(t *testing.T) {
	_, fakeProductRequest, _, mockDB := mock.NewDBSetupCreate()
	defer mockDB.Close()

	mockDB.ExpectQuery("INSERT INTO product (.+)").WithArgs(
		fakeProductRequest.Name,
		fakeProductRequest.Price,
		fakeProductRequest.Description,
	).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

	sut := productrepository.New(mockDB)
	product, err := sut.Create(&fakeProductRequest)

	if err := mockDB.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.NotNil(t, err)
	require.Nil(t, product)
}

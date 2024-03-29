package dto_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stephan-lopes/golang-clean-arch/core/dto"
	"github.com/stretchr/testify/require"
)

func TestFromJSONCreateProductRequest(t *testing.T) {
	fakeItem := dto.CreateProductRequest{}
	faker.FakeData(&fakeItem)

	json, err := json.Marshal(fakeItem)
	require.Nil(t, err)

	itemRequest, err := dto.FromJSONCreateProductRequest(strings.NewReader(string(json)))

	require.Nil(t, err)
	require.Equal(t, itemRequest.Name, fakeItem.Name)
	require.Equal(t, itemRequest.Price, fakeItem.Price)
	require.Equal(t, itemRequest.Description, fakeItem.Description)
}

func TestFromJSONCreateProductRequest_JSONDecodeError(t *testing.T) {
	itemRequest, err := dto.FromJSONCreateProductRequest(strings.NewReader("{"))

	require.NotNil(t, err)
	require.Nil(t, itemRequest)
}

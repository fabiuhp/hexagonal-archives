package application_test

import (
	"hexagonal/application"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Produto de Teste"
	product.Status = application.DISABLED
	product.Price = 10

	error := product.Enable()
	require.Nil(t, error)

	product.Price = 0
	error = product.Enable()
	require.Equal(t, "o preco deve ser maior que zero", error.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Produto de Teste"
	product.Status = application.ENABLED
	product.Price = 0

	error := product.Disable()
	require.Nil(t, error)

	product.Price = 10
	error = product.Disable()
	require.Equal(t, "o preco deve ser igual a zero", error.Error())
}

func TestIsValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Produto de Teste"
	product.Status = application.DISABLED
	product.Price = 10
	product.Id = uuid.NewV4().String()

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "o status do produto é inválido", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "o preço do produto é inválido", err.Error())

}

package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T) Produto {
	arg := CreateProdutosParams{
		Nome:      "Celular",
		Descricao: "Smartphone",
		Preco:     1299,
	}

	produto, err := testQueries.CreateProdutos(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, produto)
	require.Equal(t, arg.Nome, produto.Nome)
	require.Equal(t, arg.Descricao, produto.Descricao)
	require.Equal(t, arg.Preco, produto.Preco)
	require.NotZero(t, produto.ID)
	require.NotZero(t, produto.CriadoEm)

	return produto
}

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}

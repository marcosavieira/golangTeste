-- name: CreateProdutos :one
INSERT INTO produtos (
    nome,
    descricao,
    preco
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetProduto :one
SELECT * FROM produtos
WHERE id = $1;

-- name: ListProdutos :many
SELECT * FROM produtos
ORDER BY id;

-- name: UpdateProduto :exec
UPDATE produtos SET nome = $2,
descricao = $3,
preco = $4
WHERE id = $1;

-- name: DeleteProduto :exec
DELETE FROM produtos WHERE id = $1;
-- name: CreateTransfer :one
INSERT INTO transfers (
  amount,
  from_acount_id,
  to_account_id
) VALUES (
  $1, $2,$3
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfer :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTransfer :one
UPDATE transfers
  set amount = $2,from_acount_id=$3,to_account_id=4
WHERE id = $1
RETURNING *;


-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;
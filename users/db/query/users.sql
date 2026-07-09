-- name: CreateUser :one
INSERT INTO users (status, username, phone_number, password_hash, identity_document)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;
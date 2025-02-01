-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUsersByID :one
SELECT * FROM users WHERE id = ?;

-- name: GetUsersByIDWithSpecificField :one
SELECT id, name FROM users WHERE id = ?;

-- name: GetUserLastID :one
SELECT id FROM users ORDER BY id DESC LIMIT 1;

-- name: InsertUser :exec
INSERT INTO users (id, name, email_address) VALUES (?, ?, ?);

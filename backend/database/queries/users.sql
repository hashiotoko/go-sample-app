-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUsersByID :one
SELECT * FROM users WHERE id = ?;

-- name: GetUsersByIDWithSpecificField :one
SELECT id, name FROM users WHERE id = ?;

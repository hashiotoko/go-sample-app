// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package sqlc

import (
	"context"
)

const getUsers = `-- name: GetUsers :many
SELECT id, name, created_at, updated_at FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.query(ctx, q.getUsersStmt, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsersByID = `-- name: GetUsersByID :one
SELECT id, name, created_at, updated_at FROM users WHERE id = ?
`

func (q *Queries) GetUsersByID(ctx context.Context, id string) (User, error) {
	row := q.queryRow(ctx, q.getUsersByIDStmt, getUsersByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsersByIDWithSpecificField = `-- name: GetUsersByIDWithSpecificField :one
SELECT id, name FROM users WHERE id = ?
`

type GetUsersByIDWithSpecificFieldRow struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (q *Queries) GetUsersByIDWithSpecificField(ctx context.Context, id string) (GetUsersByIDWithSpecificFieldRow, error) {
	row := q.queryRow(ctx, q.getUsersByIDWithSpecificFieldStmt, getUsersByIDWithSpecificField, id)
	var i GetUsersByIDWithSpecificFieldRow
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

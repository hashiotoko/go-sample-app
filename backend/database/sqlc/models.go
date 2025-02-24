// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"time"
)

type User struct {
	// ユーザーID
	ID string `db:"id" json:"id"`
	// ユーザー名前
	Name string `db:"name" json:"name"`
	// 作成日時
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新日時
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	// ユーザーのメールアドレス
	EmailAddress string `db:"email_address" json:"email_address"`
}

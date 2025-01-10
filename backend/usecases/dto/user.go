package dto

import "time"

type User struct {
	ID        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

package models

import "time"

type BaseModel struct {
	ID        int64     `json:"id" db:"id"`
	CreatedAt time.Time `json:"" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
	// DeleteAt time.Time `json:"-" db:"delete_at"`
}

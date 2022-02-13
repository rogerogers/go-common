package model

import "time"

type Model struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PkModel struct {
	ID int `json:"id"`
	Model
}

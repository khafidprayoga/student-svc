package model

import (
	"time"
)

type Student struct {
	Id          string     `json:"id"`
	FullName    string     `json:"full_name"`
	Address     string     `json:"address"`
	Hobbies     []string   `json:"hobbies"`
	BirthDate   *time.Time `json:"birth_date"`
	Gender      int32      `json:"gender"`
	Email       string     `json:"email"`
	Nationality int32      `json:"nationality"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

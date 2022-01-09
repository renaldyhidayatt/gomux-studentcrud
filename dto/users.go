package dto

import "time"

type Users struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Hobby     string    `json:"hobby"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

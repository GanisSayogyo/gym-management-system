package model

import "time"

type Trainer struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	Specialization *string   `json:"specialization"`
	Phone          *string   `json:"phone"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

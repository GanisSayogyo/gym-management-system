package model

import "time"

type MembershipPlan struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Price        int64     `json:"price"`
	DurationDays int       `json:"duration_days"`
	Description  *string   `json:"description"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

package model

import "time"

type Membership struct {
	ID        string    `json:"id"`
	MemberID  string    `json:"member_id"`
	PlanID    string    `json:"plan_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

package model

import "time"

type TrainingSchedule struct {
	ID        string    `json:"id"`
	MemberID  string    `json:"member_id"`
	TrainerID string    `json:"trainer_id"`
	StartAt   time.Time `json:"start_at"`
	EndAt     time.Time `json:"end_at"`
	Status    string    `json:"status"`
	Notes     *string   `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

package model

import "time"

type Attendance struct {
	ID        string     `json:"id"`
	MemberID  string     `json:"member_id"`
	CheckIn   time.Time  `json:"check_in"`
	CheckOut  *time.Time `json:"check_out"`
	CreatedAt time.Time  `json:"created_at"`
}

package model

import "time"

type Payment struct {
	ID            string     `json:"id"`
	MembershipID  string     `json:"membership_id"`
	Amount        int64      `json:"amount"`
	PaymentMethod string     `json:"payment_method"`
	Status        string     `json:"status"`
	PaidAt        *time.Time `json:"paid_at"`
	CreatedAt     time.Time  `json:"created_at"`
}

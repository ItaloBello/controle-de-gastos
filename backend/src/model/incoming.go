package model

import "time"

type Incoming struct {
	Id          int       `db:"id" json:"id"`
	Value       float64   `db:"value" json:"value"`
	Description *string   `db:"description" json:"description"`
	IncomeDate  time.Time `db:"income_date" json:"incomeDate"`
	UserId      int       `db:"user_id" json:"userId"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
}

type IncomingCreateRequest struct {
	Value       float64   `json:"value" binding:"required"`
	Description *string   `json:"description"`
	IncomeDate  time.Time `json:"incomeDate"`
	UserId      int       `json:"userId" binding:"required"`
}

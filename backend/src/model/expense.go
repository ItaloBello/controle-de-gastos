package model

import "time"

type Expense struct {
	Id          int        `db:"id" json:"id"`
	Value       float64    `db:"value" json:"value"`
	Description *string    `db:"description" json:"description,omitempty"`
	ExpenseDate *time.Time `db:"expense_date" json:"expenseDate,omitempty"`
	IsFixed     bool       `db:"is_fixed" json:"isFixed"`
	IsEssential bool       `db:"is_essential" json:"isEssential"`
	IsPaid      bool       `db:"is_paid" json:"isPaid"`
	CategoryId  *int       `db:"category_id" json:"categoryId,omitempty"`
	UserId      int        `db:"user_id" json:"userId"`
	CreatedAt   time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updatedAt"`
}

type ExpenseCreateRequest struct {
	Value       float64 `db:"value" json:"value"`
	Description *string `db:"description" json:"description,omitempty"`
	ExpenseDate *string `db:"expense_date" json:"expenseDate,omitempty"`
	IsFixed     bool       `db:"is_fixed" json:"isFixed"`
	IsEssential bool       `db:"is_essential" json:"isEssential"`
	IsPaid      bool       `db:"is_paid" json:"isPaid"`
	CategoryId  *int    `db:"category_id" json:"categoryId,omitempty"`
	UserId      int     `db:"user_id" json:"userId"`
}

package model

type Expense struct{
	Id int `db:"id" json:"id"`
	Value float64 `db:"value" json:"value"`
	Description *string `db:"description" json:"description,omitempty"`
	ExpenseDate *string `db:"expense_date" json:"expenseDate,omitempty"`
	CategoryId *int `db:"category_id" json:"categoryId,omitempty"`
	UserId int `db:"user_id" json:"userId"`
	CreatedAt string `db:"created_at" json:"createdAt"`
	UpdatedAt string `db:"updated_at" json:"updatedAt"`
}

type ExpenseCreateRequest struct{
	Value float64 `db:"value" json:"value"`
	Description *string `db:"description" json:"description,omitempty"`
	ExpenseDate *string `db:"expense_date" json:"expenseDate,omitempty"`
	CategoryId *int `db:"category_id" json:"categoryId,omitempty"`
	UserId int `db:"user_id" json:"userId"`
}
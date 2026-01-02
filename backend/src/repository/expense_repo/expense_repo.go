package expense_repo

import (
	"controle-de-gastos/src/model"

	"github.com/jmoiron/sqlx"
)

type expenseRepo struct {
	db *sqlx.DB
}

type ExpenseRepo interface {
	GetAll() ([]model.Expense, error)
	GetById(id int) (*model.Expense, error)
	GetByUserId(userId int) ([]model.Expense, error)

	Create(expense model.ExpenseCreateRequest) (int, error)
	Update(expense model.Expense) error
	Delete(id int) error
}

func NewExpenseRepo(db *sqlx.DB) ExpenseRepo {
	return &expenseRepo{db: db}
}




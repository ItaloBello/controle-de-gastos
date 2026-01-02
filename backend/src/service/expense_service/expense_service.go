package expense_service

import (
	"controle-de-gastos/src/model"
	"controle-de-gastos/src/repository/expense_repo"
)

type expenseService struct {
	repo expense_repo.ExpenseRepo
}

type ExpenseService interface {
	GetAllExpenses() ([]model.Expense, error)
	GetExpenseById(id int) (*model.Expense, error)
	GetExpensesByUserId(userId int) ([]model.Expense, error)

	CreateExpense(expense model.ExpenseCreateRequest) (int, error)
	UpdateExpense(expense model.Expense) error
	DeleteExpense(id int) error
}

func NewExpenseService(r expense_repo.ExpenseRepo) ExpenseService {
	return &expenseService{repo: r}
}

func (s *expenseService) GetAllExpenses() ([]model.Expense, error){
	return s.repo.GetAll()
}

func (s *expenseService) GetExpenseById(id int) (*model.Expense, error){
	return s.repo.GetById(id)
}

func (s *expenseService) GetExpensesByUserId(userId int) ([]model.Expense, error){
	return s.repo.GetByUserId(userId)
}

func (s *expenseService) CreateExpense(expense model.ExpenseCreateRequest) (int, error){
	return s.repo.Create(expense)
}

func (s *expenseService) UpdateExpense(expense model.Expense) error{
	return s.repo.Update(expense)
}

func (s *expenseService) DeleteExpense(id int) error{
	return s.repo.Delete(id)
}
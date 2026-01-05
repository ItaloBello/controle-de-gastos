package expense_repo

import "controle-de-gastos/src/model"

func (r *expenseRepo) GetAll() ([]model.Expense, error) {
	const query = `SELECT id, value, description, expense_date, category_id, user_id, created_at, updated_at FROM expenses;`

	var expenses []model.Expense
	err := r.db.Select(&expenses, query)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *expenseRepo) GetById(id int) (*model.Expense, error) {
	const query = `SELECT id, value, description, expense_date, category_id, user_id, created_at, updated_at FROM expenses 
					WHERE id = $1;`

	var expense model.Expense
	err := r.db.Get(&expense, query, id)
	if err != nil {
		return nil, err
	}

	return &expense, nil
}

func (r *expenseRepo) GetByUserId(userId int) ([]model.Expense, error) {
	const query = `SELECT id, value, description, expense_date, category_id, user_id, created_at, updated_at FROM expenses 
						WHERE user_id = $1;`

	var expenses []model.Expense
	err := r.db.Select(&expenses, query, userId)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *expenseRepo) GetFixedByUserId(userId int) ([]model.Expense, error) {
	const query = `SELECT id, value, description, expense_date, is_fixed, is_essential, is_paid, category_id, user_id, created_at, updated_at FROM expenses 
						WHERE user_id = $1 AND is_fixed = TRUE;`
	var expenses []model.Expense
	err := r.db.Select(&expenses, query, userId)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *expenseRepo) GetNotFixedByUserId(userId int) ([]model.Expense, error) {
	const query = `SELECT id, value, description, expense_date, is_fixed, is_essential, is_paid, category_id, user_id, created_at, updated_at FROM expenses 
						WHERE user_id = $1 AND is_fixed = FALSE;`
	var expenses []model.Expense
	err := r.db.Select(&expenses, query, userId)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

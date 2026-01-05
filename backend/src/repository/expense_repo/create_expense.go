package expense_repo

import "controle-de-gastos/src/model"

func (r *expenseRepo) Create(expense model.ExpenseCreateRequest) (int, error) {
	const query = `INSERT INTO expenses (value, description, expense_date, is_fixed, is_essential, is_paid, category_id, user_id)
					VALUES (:value, :description, :expense_date, :is_fixed, :is_essential, :is_paid, :category_id, :user_id) RETURNING id;`

	var id int
	err := r.db.Get(&id, query, expense)

	return id, err
}

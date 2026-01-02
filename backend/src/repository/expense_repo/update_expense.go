package expense_repo

import "controle-de-gastos/src/model"

func (r *expenseRepo) Update(expense model.Expense) error {
	const query = `UPDATE expenses SET
					value = :value,
					description = :description,
					expense_date = :expense_date,
					category_id = :category_id,
					user_id = :user_id,
					updated_at = CURRENT_TIMESTAMP
					WHERE id = :id;`

	_, err := r.db.NamedExec(query, expense)

	return err
}

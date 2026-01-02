package expense_repo


func (r *expenseRepo) Delete(id int) error {
	const query = `DELETE FROM expenses WHERE id = $1;`

	_,err := r.db.Exec(query, id)

	return err
}

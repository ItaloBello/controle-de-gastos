package user_repo

func (r usuarioRepo) Delete(id int) error {
	const query = `
		DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

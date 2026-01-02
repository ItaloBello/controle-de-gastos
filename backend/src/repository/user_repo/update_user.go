package user_repo

import "controle-de-gastos/src/model"

func (r usuarioRepo) Update(user *model.User) error {
	const query = `
		UPDATE users SET 
			name = :name,
			email = :email,
			hash_pass = :hash_pass,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = :id`

	_, err := r.db.NamedExec(query, user)
	
	return err
}

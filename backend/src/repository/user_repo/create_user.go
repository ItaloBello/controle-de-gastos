package user_repo

import "controle-de-gastos/src/model"

func (r usuarioRepo) Create(user *model.User) (int, error) {
	const query = `
		INSERT INTO users(name, email, hash_pass) VALUES(:name, :email, :hash_pass)
		RETURNING id;`
	var id int

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(id)
		if err != nil {
			return 0, err
		}
	}
	
	return id, nil
}

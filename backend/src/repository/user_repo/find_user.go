package user_repo

import (
	"controle-de-gastos/src/model"
	"database/sql"
	"errors"
)

func (r usuarioRepo) GetAll() ([]model.User, error) {
	const query = `
		SELECT * FROM users`
	var users []model.User

	err := r.db.Select(users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r usuarioRepo) GetByID(id int) (*model.User, error) {
	const query = `
		SELECT * FROM users 
		WHERE id = $1`

	var user *model.User

	err := r.db.Get(user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r usuarioRepo) GetByEmailAndPassword(email, hashPass string) (*model.User, error){
	const query = `SELECT * FROM users WHERE email = $1 AND hash_pass = $2;`

	var user *model.User

	err := r.db.Get(user, query, email, hashPass)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
package user_repo

import (
	"controle-de-gastos/src/model"

	"github.com/jmoiron/sqlx"
)

type usuarioRepo struct {
	db *sqlx.DB
}

type UsuarioRepo interface {
	GetAll() ([]model.User, error)
	GetByID(id int) (*model.User, error)
	Create(user *model.User) (int, error)
	Update(id int, user *model.User) error
	Delete(id int) error
}

func NewUsuarioRepo(db *sqlx.DB) UsuarioRepo {
	return &usuarioRepo{db: db}
}

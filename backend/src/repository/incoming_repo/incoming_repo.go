package incoming_repo

import (
	"controle-de-gastos/src/model"

	"github.com/jmoiron/sqlx"
)

type incomingRepo struct {
	db *sqlx.DB
}

type IncomingRepo interface {
	GetAll() ([]model.Incoming, error)
	GetById(id int) (*model.Incoming, error)

	Create(incoming model.Incoming) (int, error)
	Update(incoming model.Incoming) error
	Delete(id int) error
}

func NewIncomingRepo(db *sqlx.DB) IncomingRepo {
	return &incomingRepo{db: db}
}

func (r *incomingRepo) GetAll() ([]model.Incoming, error) {
	const query = `SELECT id, value, description, income_date, user_id, created_at, updated_at FROM incomings`
	var incomings []model.Incoming
	err := r.db.Select(&incomings, query)
	if err != nil {
		return nil, err
	}
	return incomings, nil
}

func (r *incomingRepo) GetById(id int) (*model.Incoming, error) {
	const query = `SELECT id, value, description, income_date, user_id, created_at, updated_at FROM incomings WHERE id = $1`
	var incoming model.Incoming
	err := r.db.Get(&incoming, query, id)
	if err != nil {
		return nil, err
	}
	return &incoming, nil
}

func (r *incomingRepo) Create(incoming model.Incoming) (int, error) {
	const query = `INSERT INTO incomings (value, description, income_date, user_id, created_at, updated_at) 
					VALUES (:value, :description, :income_date, :user_id, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)  RETURNING id`
	var id int
	err := r.db.Get(&id, query, incoming)
	return id, err
}

func (r *incomingRepo) Update(incoming model.Incoming) error {
	const query = `UPDATE incomings SET 
					value = :value, 
					description = :description, 
					income_date = :income_date, 
					updated_at = CURRENT_TIMESTAMP
					WHERE id = :id`
	_, err := r.db.NamedExec(query, incoming)
	return err
}

func (r *incomingRepo) Delete(id int) error {
	const query = `DELETE FROM incomings WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
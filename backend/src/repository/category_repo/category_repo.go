package category_repo

import (
	"controle-de-gastos/src/model"

	"github.com/jmoiron/sqlx"
)

type categoryRepo struct {
	db *sqlx.DB
}

type CategoryRepo interface {
	GetAll() ([]model.Category, error)
	GetById(id int) (*model.Category, error)

	Create(category model.Category) (int, error)
	Update(category model.Category) error
	Delete(id int) error
}

func NewCategoryRepo(db *sqlx.DB) CategoryRepo {
	return &categoryRepo{db: db}
}

func (r *categoryRepo) GetAll() ([]model.Category, error) {
	const query = `SELECT id, name FROM categories;`
	var categories []model.Category
	err := r.db.Select(&categories, query)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepo) GetById(id int) (*model.Category, error) {
	const query = `SELECT id, name FROM categories WHERE id = $1;`
	var category model.Category
	err := r.db.Get(&category, query, id)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepo) Create(category model.Category) (int, error) {
	const query = `INSERT INTO categories (name) VALUES (:name) RETURNING id;`

	var id int
	err := r.db.Get(&id, query, category)
	return id, err
}

func (r *categoryRepo) Update(category model.Category) error {
	const query = `UPDATE categories SET name = :name WHERE id = :id;`

	_, err := r.db.NamedExec(query, category)
	return err
}

func (r *categoryRepo) Delete(id int) error {
	const query = `DELETE FROM categories WHERE id = $1;`

	_, err := r.db.Exec(query, id)
	return err
}

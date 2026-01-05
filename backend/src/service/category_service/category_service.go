package category_service

import (
	"controle-de-gastos/src/model"
	"controle-de-gastos/src/repository/category_repo"
)

type categoryService struct {
	categoryRepo category_repo.CategoryRepo
}

type CategoryService interface {
	GetAll() ([]model.Category, error)
	GetById(id int) (*model.Category, error)
	Create(category model.Category) (int, error)
	Update(category model.Category) error
	Delete(id int) error
}

func NewCategoryService(categoryRepo category_repo.CategoryRepo) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (s *categoryService) GetAll() ([]model.Category, error) {
	return s.categoryRepo.GetAll()
}

func (s *categoryService) GetById(id int) (*model.Category, error) {
	return s.categoryRepo.GetById(id)
}

func (s *categoryService) Create(category model.Category) (int, error) {
	return s.categoryRepo.Create(category)
}

func (s *categoryService) Update(category model.Category) error {
	return s.categoryRepo.Update(category)
}

func (s *categoryService) Delete(id int) error {
	return s.categoryRepo.Delete(id)
}
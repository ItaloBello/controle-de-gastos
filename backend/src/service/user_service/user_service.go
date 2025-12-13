package user_service

import (
	"controle-de-gastos/src/model"
	"controle-de-gastos/src/repository/user_repo"
)

type userService struct {
	repo user_repo.UsuarioRepo
}

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserById(id int) (*model.User, error)
	CreateUser(user *model.User) (int, error)
	UpdateUser(id int, user *model.User) error
	DeleteUser(id int) error
}

func NewUserService(r user_repo.UsuarioRepo) UserService {
	return &userService{repo: r}
}

func (s userService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s userService) GetUserById(id int) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s userService) CreateUser(user *model.User) (int, error) {
	return s.repo.Create(user)
}

func (s userService) UpdateUser(id int, user *model.User) error {
	return s.repo.Update(id, user)
}

func (s userService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}

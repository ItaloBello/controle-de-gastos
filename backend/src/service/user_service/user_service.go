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
	CreateUser(user *model.UserCreateRequest) (int, error)
	UpdateUser(id int, user *model.User) error
	DeleteUser(id int) error

	LoginUser(userLogin *model.UserLogin) (*model.User, error)
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

func (s userService) CreateUser(userCreate *model.UserCreateRequest) (int, error) {
	user := model.User{
		Name:     userCreate.Name,
		Email:    userCreate.Email,
		HashPass: userCreate.Password,
	}
	//criptografar a senha
	user.EncryptPassword()

	return s.repo.Create(&user)
}

func (s userService) UpdateUser(id int, user *model.User) error {
	user.ID = id
	return s.repo.Update(user)
}

func (s userService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}

func (s userService) LoginUser(userLogin *model.UserLogin) (*model.User, error) {
	requestUser := model.User{
		Email:    userLogin.Email,
		HashPass: userLogin.Password,
	}
	//criptografar a senha
	requestUser.EncryptPassword()

	hashPass := userLogin.Password
	user, err := s.repo.GetByEmailAndPassword(userLogin.Email, hashPass)
	if err != nil {
		return nil, err
	}
	return user, nil
}

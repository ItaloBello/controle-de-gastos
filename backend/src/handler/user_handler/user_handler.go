package user_handler

import (
	"controle-de-gastos/src/model"
	"controle-de-gastos/src/service/user_service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user_service.UserService
}

type UserHandler interface {
	GetAllUsers(c *gin.Context)
	GetUserById(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)

	UserLogin(c *gin.Context)
}

func NewUserHandler(serv user_service.UserService) UserHandler {
	return &userHandler{service: serv}
}

func (h userHandler) GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
func (h userHandler) GetUserById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
func (h userHandler) CreateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"status": "ok"})
}
func (h userHandler) UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
func (h userHandler) DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h userHandler) UserLogin(c *gin.Context) {
	var userLogin model.UserLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error biding JSON", "error": err.Error()})
		fmt.Println("error biding JSON: " + err.Error())
		return
	}
	user, err := h.service.LoginUser(&userLogin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error in login", "error": err.Error()})
		fmt.Println("error in login: " + err.Error())
		return
	}
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"access": false})
		fmt.Printf("Forbidden login for email: %s\n", userLogin.Email)
		return
	}

	c.JSON(http.StatusOK, gin.H{"access": true, "role": "user", "name": user.Name, "email": user.Email})
}

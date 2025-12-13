package user_handler

import (
	"controle-de-gastos/src/service/user_service"
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
func (h userHandler) CreateUser(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
func (h userHandler) UpdateUser(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
func (h userHandler) DeleteUser(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

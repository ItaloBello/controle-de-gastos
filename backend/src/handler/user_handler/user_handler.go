package user_handler

import (
	"controle-de-gastos/src/model"
	"controle-de-gastos/src/service/user_service"
	"fmt"
	"net/http"
	"strconv"

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
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "get all users error", "error": err.Error()})
		fmt.Println("get all users error: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}
func (h userHandler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id must be a integer", "error": err.Error()})
		fmt.Println("id must be a integer: " + err.Error())
		return
	}

	user, err := h.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "get user by id error", "error": err.Error()})
		fmt.Println("get user by id error: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
func (h userHandler) CreateUser(c *gin.Context) {
	var user model.UserCreateRequest
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message":"error binding json in CreateUser", "error":err.Error()})
		fmt.Println("error binding json in CreateUser: "+err.Error())
		return
	}

	id, err := h.service.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "create user error", "error": err.Error()})
		fmt.Println("create user error: " + err.Error())
		return
	}

	c.JSON(http.StatusCreated, id)
}
func (h userHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id must be a integer", "error": err.Error()})
		fmt.Println("id must be a integer: " + err.Error())
		return
	}

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message":"error binding json in CreateUser", "error":err.Error()})
		fmt.Println("error binding json in CreateUser: "+err.Error())
		return
	}

	err = h.service.UpdateUser(id, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "update user error", "error": err.Error()})
		fmt.Println("update user error: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update successfully"})
}
func (h userHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id must be a integer", "error": err.Error()})
		fmt.Println("id must be a integer: " + err.Error())
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "update user error", "error": err.Error()})
		fmt.Println("update user error: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete successfully"})
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

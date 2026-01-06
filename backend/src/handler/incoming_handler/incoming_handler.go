package incoming_handler

import (
	"controle-de-gastos/src/model"
	"controle-de-gastos/src/service/incoming_service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type incomingHandler struct {
	service incoming_service.IncomingService
}

type IncomingHandler interface {
	GetAllIncomes(c *gin.Context)
	GetIncomingById(c *gin.Context)
	CreateIncoming(c *gin.Context)
	UpdateIncoming(c *gin.Context)
	DeleteIncoming(c *gin.Context)
}

func NewIncomingHandler(serv incoming_service.IncomingService) IncomingHandler {
	return &incomingHandler{service: serv}
}

func (h *incomingHandler) GetAllIncomes(c *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "get all incomes error", "error": err.Error()})
		fmt.Println("get all users error: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *incomingHandler) GetIncomingById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id must be a integer", "error": err.Error()})
		fmt.Println("id must be a integer: " + err.Error())
		return
	}

	incoming, err := h.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "get incoming by id error", "error": err.Error()})
		fmt.Println("get incoming by id error: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, incoming)
}

func (h *incomingHandler) CreateIncoming(c *gin.Context) {
	var incoming model.IncomingCreateRequest
	if err := c.ShouldBindJSON(&incoming); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message":"error binding json in CreateIncoming", "error":err.Error()})
		fmt.Println("error binding json in CreateIncoming: "+err.Error())
		return
	}

	id, err := h.service.Create(incoming)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "create incoming error", "error": err.Error()})
		fmt.Println("create incoming error: " + err.Error())
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (h *incomingHandler) UpdateIncoming(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id must be a integer", "error": err.Error()})
		fmt.Println("id must be a integer: " + err.Error())
		return
	}

	var incoming model.Incoming
	if err := c.ShouldBindJSON(&incoming); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message":"error binding json in UpdateIncoming", "error":err.Error()})
		fmt.Println("error binding json in UpdateIncoming: "+err.Error())
		return
	}

	err = h.service.Update(id, incoming)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "update incoming error", "error": err.Error()})
		fmt.Println("update incoming error: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update successfully"})
}

func (h *incomingHandler) DeleteIncoming(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id must be a integer", "error": err.Error()})
		fmt.Println("id must be a integer: " + err.Error())
		return
	}
	
	err = h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "update user error", "error": err.Error()})
		fmt.Println("update user error: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete successfully"})
}
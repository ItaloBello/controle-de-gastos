package expense_handler

import (
	"controle-de-gastos/src/model"
	"controle-de-gastos/src/service/expense_service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type expenseHandler struct {
	service expense_service.ExpenseService
}

type ExpenseHandler interface {
	GetAllExpenses(c *gin.Context)
	GetExpenseById(c *gin.Context)
	GetExpensesByUserId(c *gin.Context)

	CreateExpense(c *gin.Context)
	UpdateExpense(c *gin.Context)
	DeleteExpense(c *gin.Context)
}

func NewExpenseHandler(serv expense_service.ExpenseService) ExpenseHandler {
	return &expenseHandler{service: serv}
}

func (h *expenseHandler) GetAllExpenses(c *gin.Context) {
	expenses, err := h.service.GetAllExpenses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "get all expenses error", "error": err.Error()})
		fmt.Println("get all expenses error: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func (h *expenseHandler) GetExpenseById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id must be a integer", "error": err.Error()})
		fmt.Println("id must be a integer: " + err.Error())
		return
	}

	expense, err := h.service.GetExpenseById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "get expense by id error", "error": err.Error()})
		fmt.Println("get expense by id error: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, expense)
}

func (h *expenseHandler) GetExpensesByUserId(c *gin.Context){
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "userId must be a integer", "error": err.Error()})
		fmt.Println("userId must be a integer: " + err.Error())
		return
	}

	expenses, err := h.service.GetExpensesByUserId(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "get expenses by userId error", "error": err.Error()})
		fmt.Println("get expenses by userId error: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func (h *expenseHandler) CreateExpense(c *gin.Context){
	var expense model.ExpenseCreateRequest
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"error binding json in CreateExpense", "error":err.Error()})
		fmt.Println("error binding json in CreateExpense: "+err.Error())
		return
	}

	id, err := h.service.CreateExpense(expense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "create expense error", "error": err.Error()})
		fmt.Println("create expense error: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *expenseHandler) UpdateExpense(c *gin.Context){
	var expense model.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"error binding json in UpdateExpense", "error":err.Error()})
		fmt.Println("error binding json in UpdateExpense: "+err.Error())
		return
	}

	err := h.service.UpdateExpense(expense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "update expense error", "error": err.Error()})
		fmt.Println("update expense error: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"update successfully"})
}

func (h *expenseHandler) DeleteExpense(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id must be a integer", "error": err.Error()})
		fmt.Println("id must be a integer: " + err.Error())
		return
	}
	err = h.service.DeleteExpense(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "delete expense error", "error": err.Error()})
		fmt.Println("delete expense error: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"delete successfully"})
}
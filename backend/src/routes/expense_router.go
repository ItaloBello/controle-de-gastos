package routes

import (
	"controle-de-gastos/src/handler/expense_handler"

	"github.com/gin-gonic/gin"
)

func SetupExpenseRoutes(r *gin.RouterGroup, h expense_handler.ExpenseHandler) {
	user := r.Group("/expense")
	{
		user.GET("/", h.GetAllExpenses)
		user.GET("/:id", h.GetExpenseById)
		user.GET("/user/:userId", h.GetExpensesByUserId)
		user.GET("/fixed/:userId", h.GetFixedExpensesByUserId) // Vou realmente usar
		user.GET("/not-fixed/:userId", h.GetNotFixedExpensesByUserId) // Vou realmente usar

		user.POST("/", h.CreateExpense)
		user.PUT("/:id", h.UpdateExpense)
		user.DELETE("/:id", h.DeleteExpense)
	}
}

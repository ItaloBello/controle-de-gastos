package routes

import (
	"controle-de-gastos/src/handler/category_handler"

	"github.com/gin-gonic/gin"
)

func SetupCategoryRoutes(r *gin.RouterGroup, h category_handler.CategoryHandler) {
	user := r.Group("/category")
	{
		user.GET("/", h.GetAllCategories)
		user.GET("/:id", h.GetCategoryById)
		user.POST("/", h.CreateCategory)
		user.PUT("/:id", h.UpdateCategory)
		user.DELETE("/:id", h.DeleteCategory)
	}
}

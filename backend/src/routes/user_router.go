package routes

import (
	"controle-de-gastos/src/handler/user_handler"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup, h user_handler.UserHandler){
	user := r.Group("/user")
	{
		user.GET("/", h.GetAllUsers)
		user.GET("/:id", h.GetUserById)
		user.POST("/", h.CreateUser)
		user.PUT("/:id", h.UpdateUser)
		user.DELETE("/:id", h.DeleteUser)
	}
}
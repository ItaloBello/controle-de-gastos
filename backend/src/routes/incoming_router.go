package routes

import (

	"controle-de-gastos/src/handler/incoming_handler"

	"github.com/gin-gonic/gin"
)

func SetupIncomingRoutes(r *gin.RouterGroup, h incoming_handler.IncomingHandler) {
	user := r.Group("/incoming")
	{
		user.GET("/", h.GetAllIncomes)
		user.GET("/:id", h.GetIncomingById)
		user.POST("/", h.CreateIncoming)
		user.PUT("/:id", h.UpdateIncoming)
		user.DELETE("/:id", h.DeleteIncoming)
	}
}

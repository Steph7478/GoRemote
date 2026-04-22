package routes

import (
	"server/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.GET("/status", handlers.Status)
	r.GET("/ws", handlers.WebSocket)

	api := r.Group("/api")
	{
		mouse := api.Group("/mouse")
		{
			mouse.POST("/move", handlers.MouseMove)
			mouse.POST("/click", handlers.MouseClick)
			mouse.POST("/scroll", handlers.MouseScroll)
		}

		key := api.Group("/key")
		{
			key.POST("/press", handlers.KeyPress)
			key.POST("/type", handlers.KeyType)
		}
	}
}

package handlers

import (
	"net/http"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "online",
		"ip":     utils.GetLocalIP(),
		"port":   "8080",
	})
}

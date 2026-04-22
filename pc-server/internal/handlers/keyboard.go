package handlers

import (
	"net/http"
	"server/internal/models"

	"github.com/gin-gonic/gin"
)

func KeyPress(c *gin.Context) {
	var cmd models.Command
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	robot.Press(cmd.Key)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func KeyType(c *gin.Context) {
	var cmd models.Command
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	robot.Type(cmd.Text)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

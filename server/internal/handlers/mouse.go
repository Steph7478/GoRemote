package handlers

import (
	"net/http"
	"server/internal/controller"
	"server/internal/models"

	"github.com/gin-gonic/gin"
)

var robot = controller.New()

func MouseMove(c *gin.Context) {
	var cmd models.Command
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	robot.Move(cmd.X, cmd.Y)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func MouseClick(c *gin.Context) {
	robot.LeftClick()
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func MouseScroll(c *gin.Context) {
	var cmd models.Command
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	robot.Scroll(int(cmd.X), int(cmd.Y))
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

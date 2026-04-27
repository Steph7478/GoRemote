package handlers

import (
	"net/http"
	"server/internal/controller"
	"server/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var robot = controller.New()

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

var wsHandlers = map[string]func(models.WSMessage){
	"move":        func(m models.WSMessage) { robot.Move(m.X, m.Y) },
	"left_click":  func(m models.WSMessage) { robot.LeftClick() },
	"right_click": func(m models.WSMessage) { robot.RightClick() },
	"scroll":      func(m models.WSMessage) { robot.Scroll(int(m.X), int(m.Y)) },
	"down":        func(m models.WSMessage) { robot.MouseDown() },
	"up":          func(m models.WSMessage) { robot.MouseUp() },
	"key":         func(m models.WSMessage) { robot.Press(m.Key) },
	"type":        func(m models.WSMessage) { robot.Type(m.Text) },
}

func WebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	var m models.WSMessage
	for conn.ReadJSON(&m) == nil {
		if fn, ok := wsHandlers[m.Event]; ok {
			fn(m)
		}
	}
}

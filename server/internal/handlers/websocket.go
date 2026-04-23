package handlers

import (
	"net/http"
	"server/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

var wsHandlers = map[string]func(models.WSMessage){
	"move": func(m models.WSMessage) { robot.Move(m.X, m.Y) },

	"click": func(m models.WSMessage) { robot.Click() },

	"key": func(m models.WSMessage) { robot.Press(m.Key) },

	"type": func(m models.WSMessage) { robot.Type(m.Text) },
}

func WebSocket(c *gin.Context) {
	conn, _ := upgrader.Upgrade(c.Writer, c.Request, nil)
	defer conn.Close()

	var m models.WSMessage
	for conn.ReadJSON(&m) == nil {
		if fn, ok := wsHandlers[m.Event]; ok {
			fn(m)
		}
	}
}

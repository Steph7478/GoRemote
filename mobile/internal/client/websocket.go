package client

import (
	"mobile/internal/models"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn        *websocket.Conn
	addr        string
	mu          sync.Mutex
	sensitivity float64
}

func NewClient(ip string) *Client {
	return &Client{
		addr:        ip,
		sensitivity: 1.0,
	}
}

func (c *Client) SetSensitivity(value float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sensitivity = value
}

func (c *Client) Connect() error {
	u := url.URL{Scheme: "ws", Host: c.addr + ":8080", Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *Client) Send(msg models.WSMessage) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if msg.Event == "move" && c.sensitivity != 1.0 {
		msg.X = msg.X * c.sensitivity
		msg.Y = msg.Y * c.sensitivity
	}

	return c.conn.WriteJSON(msg)
}

func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

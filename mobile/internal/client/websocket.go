package client

import (
	"errors"
	"mobile/internal/models"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/schollz/peerdiscovery"
)

type Client struct {
	conn        *websocket.Conn
	addr        string
	sensitivity float64
	mu          sync.Mutex
}

func NewClient(ip string) *Client {
	return &Client{
		addr:        ip,
		sensitivity: 1.0,
	}
}

func (c *Client) SetSensitivity(v float64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.sensitivity = v
}

func (c *Client) Connect() error {
	u := url.URL{
		Scheme: "ws",
		Host:   c.addr + ":8080",
		Path:   "/ws",
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	c.mu.Lock()
	c.conn = conn
	c.mu.Unlock()

	return nil
}

func (c *Client) Send(msg models.WSMessage) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn == nil {
		return errors.New("websocket disconnected")
	}

	if msg.Event == "move" {
		msg.X *= c.sensitivity
		msg.Y *= c.sensitivity
	}

	return c.conn.WriteJSON(msg)
}

func (c *Client) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		_ = c.conn.Close()
		c.conn = nil
	}
}

func DiscoverServer() (string, error) {
	discoveries, err := peerdiscovery.Discover(peerdiscovery.Settings{
		Limit:     1,
		Delay:     time.Second,
		TimeLimit: 5 * time.Second,
		AllowSelf: false,
	})
	if err != nil {
		return "", err
	}

	for _, d := range discoveries {
		if string(d.Payload) == "RemoteControl" && d.Address != "" {
			return d.Address, nil
		}
	}

	return "", nil
}

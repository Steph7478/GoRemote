package client

import (
	"mobile/internal/models"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/schollz/peerdiscovery"
)

type Client struct {
	conn        *websocket.Conn
	addr        string
	sensitivity float64
}

func NewClient(ip string) *Client {
	return &Client{addr: ip, sensitivity: 1.0}
}

func (c *Client) SetSensitivity(v float64) {
	c.sensitivity = v
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
	if msg.Event == "move" && c.sensitivity != 1.0 {
		msg.X *= c.sensitivity
		msg.Y *= c.sensitivity
	}
	return c.conn.WriteJSON(msg)
}

func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func DiscoverServer() (string, error) {
	discoveries, err := peerdiscovery.Discover(peerdiscovery.Settings{
		Limit:     1,
		Delay:     time.Second * 1,
		TimeLimit: time.Second * 5,
		AllowSelf: true,
	})
	if err != nil {
		return "", err
	}

	for _, d := range discoveries {
		payload := string(d.Payload)

		if strings.Contains(payload, "RemoteControl:") {
			parts := strings.SplitN(payload, ":", 2)
			if len(parts) == 2 {
				return parts[1], nil
			}
		}

		if d.Address != "" {
			return d.Address, nil
		}
	}

	return "", nil
}
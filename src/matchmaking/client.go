package matchmaking

import (
	"fmt"
	"log"
	"time"

	"github.com/0110101001110011/blade2/src/templates"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

const HeartbeatPeriod = 1 * time.Second

var birthTime = time.Now().Unix()

type Client struct {
	Connection *websocket.Conn
	ID         string
	SendQueue  chan []byte
}

func NewClient(c *websocket.Conn) Client {
	id := uuid.Must(uuid.NewV4()).String()
	return Client{c, id, make(chan []byte)}
}

func (c *Client) run() {
	go func() {
		for {
			message := <-c.SendQueue
			err := c.Connection.WriteJSON(message)

			if err != nil {
				// TODO handle error
			}
		}
	}()

	go func() {
		for {
			_, message, err := c.Connection.ReadMessage()

			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("error: %v", err)
				}
				break
			}

			fmt.Printf("Received a message: [%s]\n", message)
		}
	}()

	go func() {
		for {
			time.Sleep(HeartbeatPeriod)

			c.sendMessage(templates.Make(templates.Heartbeat{Uptime: time.Now().Unix() - birthTime}))
		}
	}()
}

func (c *Client) sendMessage(msg []byte) {
	c.SendQueue <- msg
}

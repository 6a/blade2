package matchmaking

import (
	"fmt"
	"log"
	"time"

	"github.com/0110101001110011/blade2/src/templates"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

const heartbeatPeriod = 1 * time.Second
const heartbeatTimeout = heartbeatPeriod * 2
const dropDelay = 2 * time.Second

var birthTime = time.Now().Unix()
var ack = []byte("LS1BQ0stLQ==")

// Client is a container for a websocket connection with a client, along with some metadata
type Client struct {
	Connection      *websocket.Conn
	ID              string
	SendQueue       chan []byte
	LastMessageTime time.Time
	DropOnNextSend  bool
}

// NewClient creates a new Client object with an associated websocket connection
func NewClient(c *websocket.Conn) Client {
	id := uuid.Must(uuid.NewV4()).String()
	return Client{c, id, make(chan []byte), time.Now(), false}
}

func (c *Client) activate() {
	c.LastMessageTime = time.Now()
	go sendPump(c)

	go receivePump(c)

	go heartbeat(c)
}

func sendPump(c *Client) {
	for {
		message := <-c.SendQueue
		err := c.Connection.WriteJSON(message)

		if err != nil {
			// TODO handle error

		}

		if c.DropOnNextSend {
			err := c.Connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}

			time.Sleep(dropDelay)
			c.Connection.Close()
		}
	}
}

func receivePump(c *Client) {
	for {
		_, message, err := c.Connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break // TODO handle
		}

		c.LastMessageTime = time.Now()

		if string(message) != string(ack) {
			fmt.Printf("[%s] Received a message: [%s]\n", c.ID, message)
		}
	}
}

func heartbeat(c *Client) {
	for {
		time.Sleep(heartbeatPeriod)

		c.sendMessage(templates.MakeJSON(templates.Heartbeat{Uptime: time.Now().Unix() - birthTime}))
	}
}

func (c *Client) sendMessage(msg []byte) {
	c.SendQueue <- msg
}

// IsAlive returns true if the time since most recent client response is less than the heartbeat timeout value
func (c *Client) IsAlive() bool {
	return (time.Since(c.LastMessageTime) < heartbeatTimeout)
}

// Drop sends a Template to pass through to the client. After calling this function the client will be dropped gracefully if possible
func (c *Client) Drop(v interface{}) {
	c.sendMessage(templates.MakeJSON(v))
	c.DropOnNextSend = true
}

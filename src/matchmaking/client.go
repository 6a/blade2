package matchmaking

import (
	"fmt"
	"log"
	"time"

	"github.com/0110101001110011/blade2/src/templates"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

const heartbeatPeriod = 5 * time.Second
const heartbeatTimeout = heartbeatPeriod * 2
const dropDelay = 2 * time.Second

var birthTime = time.Now().Unix()
var ack = []byte("LS1IQi0t")
var heartbeatPayload = []byte("--HB--")

// Client is a container for a websocket connection with a client, along with some metadata
type Client struct {
	Connection      *websocket.Conn
	ID              string
	SendQueue       chan []byte
	Update          *templates.StateUpdate
	LastMessageTime time.Time
	DropOnNextSend  bool
}

// NewClient creates a new Client object with an associated websocket connection
func NewClient(c *websocket.Conn) Client {
	id := uuid.Must(uuid.NewV4()).String()
	return Client{c, id, make(chan []byte), nil, time.Now(), false}
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
		var err error
		if string(message) == string(heartbeatPayload) {
			err = c.Connection.WriteMessage(websocket.BinaryMessage, message)
		} else {
			err = c.Connection.WriteJSON(message)
		}

		if err != nil {
			// TODO handle
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

		c.sendHeartbeat()
	}
}

func (c *Client) sendMessage(msg []byte) {
	c.SendQueue <- msg
}

func (c *Client) sendHeartbeat() {
	c.SendQueue <- heartbeatPayload
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

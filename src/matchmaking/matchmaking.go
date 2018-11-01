package matchmaking

import (
	"fmt"
	"time"

	"github.com/0110101001110011/blade2/src/e"
	"github.com/0110101001110011/blade2/src/templates"
	"github.com/gorilla/websocket"
)

const maxChannelSize = 10240
const maxPollWait = 1000 * time.Millisecond

var matchmakingQueue = make(chan *Client, maxChannelSize)
var running = true
var initTime int64

func poll() {
	for {
		if len(matchmakingQueue) > 1 {
			c1, c2 := <-matchmakingQueue, <-matchmakingQueue

			if c1.IsAlive() && c2.IsAlive() {
				c1.sendMessage(templates.MakeJSON(templates.Information{Code: e.MatchFound, Message: ""}))
				c2.sendMessage(templates.MakeJSON(templates.Information{Code: e.MatchFound, Message: ""}))
				game := CreateGame(c1, c2)
				AddGame(&game)
			} else {
				if !c1.IsAlive() {
					c1.Drop(templates.MakeJSON(templates.Information{Code: e.OponentDroppedConnection, Message: ""}))
				} else {
					matchmakingQueue <- c1
				}

				if !c2.IsAlive() {
					c2.Drop(templates.MakeJSON(templates.Information{Code: e.OponentDroppedConnection, Message: ""}))
				} else {
					matchmakingQueue <- c2
				}
			}
		} else {
			time.Sleep(maxPollWait)

			if !running {
				break
			}
		}
	}
}

// JoinQueue creates a Client object and adds it to the matchmaking queue
func JoinQueue(c *websocket.Conn) {
	client := NewClient(c)
	fmt.Printf("Added client [%s] to the matchmaking queue\n", client.ID)
	client.activate()

	matchmakingQueue <- &client
	client.sendMessage(templates.MakeJSON(templates.Information{Code: e.Connected, Message: client.ID}))
}

// InitMatchMakingQueue initializes the matchmaking queue
func InitMatchMakingQueue() {
	initTime = time.Now().Unix()
	go poll()
}

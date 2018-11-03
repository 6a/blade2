package matchmaking

import (
	"fmt"
	"time"

	"github.com/0110101001110011/blade2/src/e"
	"github.com/0110101001110011/blade2/src/templates"
	"github.com/gorilla/websocket"
)

const maxPollWait = 1000 * time.Millisecond

var matchmakingQueue []*Client
var running = true
var initTime int64

func poll() {
	for {
		if len(matchmakingQueue) > 1 {
			pair, matchmakingQueue := matchmakingQueue[0:2], matchmakingQueue[2:]
			c1, c2 := pair[0], pair[1]
			if c1.IsAlive() && c2.IsAlive() {
				c1.sendMessage(templates.MakeJSON(templates.Information{Code: e.MatchFound, Message: "0"}))
				c2.sendMessage(templates.MakeJSON(templates.Information{Code: e.MatchFound, Message: "1"}))
				game := CreateGame(c1, c2)
				AddGame(&game)
			} else {
				if !c1.IsAlive() {
					c1.Drop(templates.MakeJSON(templates.Information{Code: e.OponentDroppedConnection, Message: ""}))
				} else {
					matchmakingQueue = append([]*Client{c1}, matchmakingQueue...)
				}

				if !c2.IsAlive() {
					c2.Drop(templates.MakeJSON(templates.Information{Code: e.OponentDroppedConnection, Message: ""}))
				} else {
					matchmakingQueue = append([]*Client{c2}, matchmakingQueue...)
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

	matchmakingQueue = append(matchmakingQueue, &client)
	client.sendMessage(templates.MakeJSON(templates.Information{Code: e.Connected, Message: client.ID}))
}

// InitMatchMakingQueue initializes the matchmaking queue
func InitMatchMakingQueue() {
	initTime = time.Now().Unix()
	go poll()
}

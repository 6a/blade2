package server

import (
	"fmt"
	"strconv"
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
		// If there are at least 2 clients in the matchmaking queue
		if len(matchmakingQueue) > 1 {
			// Find two live clients
			var clients []*Client
			queuePos := 0
			for {
				// If the current client being checked is alive, add it to the list of clients
				// Otherwise drop it and keep searching
				if matchmakingQueue[queuePos].IsAlive() {
					clients = append(clients, matchmakingQueue[queuePos])
				} else {
					matchmakingQueue[queuePos].Drop(templates.MakeJSON(templates.Information{Code: e.Drop, Message: ""}))
				}

				queuePos++

				// If two live clients are found, generate a match and add it to the game host
				if len(clients) > 1 {
					matchmakingQueue = matchmakingQueue[queuePos:]
					for index := 0; index < 2; index++ {
						clients[index].sendMessage(templates.MakeJSON(templates.Information{Code: e.MatchFound, Message: strconv.Itoa(index)}))
					}

					game := CreateGame(clients[0], clients[1])
					AddGame(&game)
					break
				}

				// If the number of clients that are needed exceeds the number of unchecked clients in the queue, exit early
				if len(matchmakingQueue[queuePos:]) < (2 - len(clients)) {
					matchmakingQueue = matchmakingQueue[queuePos:]
					break
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
